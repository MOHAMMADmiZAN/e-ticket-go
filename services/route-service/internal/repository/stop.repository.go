package repository

import (
	"context"
	"errors"
	"route-service/internal/api/dto"
	"route-service/internal/models"

	"gorm.io/gorm"
)

var (
	ErrStopNotFound = errors.New("stop not found")
	// You can define more specific errors related to database operations if needed.
)

// StopRepositoryInterface defines the methods that our StopRepository must implement.
type StopRepositoryInterface interface {
	AddStopToRoute(ctx context.Context, stop models.Stop) (*dto.StopResponse, error)
	ListAllStopsForRoute(ctx context.Context, routeID uint) ([]dto.StopResponse, error)
	UpdateStopDetails(ctx context.Context, stop models.Stop) (*dto.StopResponse, error)
	DeleteStop(ctx context.Context, routeID uint, stopID uint) error
	FindStopByID(ctx context.Context, stopID uint) (*dto.StopResponse, error)
	ReSequenceStops(ctx context.Context, routeID uint, deletedStopSequence int) error
	DeleteAndReSequenceStops(ctx context.Context, routeID uint, stopID uint) error
	GetStopBySequenceAndRouteID(ctx context.Context, sequence int, routeID uint) (*models.Stop, error)
	GetStopByID(ctx context.Context, routID uint, stopID uint) (*dto.StopResponse, error)
}

// Ensure that StopRepository implements the StopRepositoryInterface.
var _ StopRepositoryInterface = (*StopRepository)(nil)

type StopRepository struct {
	db *gorm.DB
}

func NewStopRepository(db *gorm.DB) *StopRepository {
	return &StopRepository{db: db}
}

func (repo *StopRepository) AddStopToRoute(ctx context.Context, stop models.Stop) (*dto.StopResponse, error) {
	if err := repo.db.WithContext(ctx).Create(&stop).Error; err != nil {
		return nil, err
	}

	if err := repo.db.WithContext(ctx).Preload("Route").First(&stop, stop.ID).Error; err != nil {
		return nil, err
	}

	stopResponse := mapStopToResponse(stop)
	return &stopResponse, nil
}

func (repo *StopRepository) ListAllStopsForRoute(ctx context.Context, routeID uint) ([]dto.StopResponse, error) {
	var stops []models.Stop
	if err := repo.db.WithContext(ctx).
		Where("route_id = ?", routeID).
		Order("sequence asc").
		Find(&stops).Error; err != nil {
		return nil, err
	}

	var stopResponses []dto.StopResponse
	for _, stop := range stops {
		stopResponses = append(stopResponses, mapStopToResponse(stop))
	}

	return stopResponses, nil
}

func (repo *StopRepository) UpdateStopDetails(ctx context.Context, stop models.Stop) (*dto.StopResponse, error) {
	if err := repo.db.WithContext(ctx).Save(&stop).Error; err != nil {
		return nil, err
	}

	if err := repo.db.WithContext(ctx).Preload("Route").First(&stop, stop.ID).Error; err != nil {
		return nil, err
	}

	stopResponse := mapStopToResponse(stop)
	return &stopResponse, nil
}

func (repo *StopRepository) DeleteStop(ctx context.Context, routeID uint, stopID uint) error {
	if err := repo.db.WithContext(ctx).Where("id = ? AND route_id = ?", stopID, routeID).Delete(&models.Stop{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *StopRepository) FindStopByID(ctx context.Context, stopID uint) (*dto.StopResponse, error) {
	var stop models.Stop
	if err := repo.db.WithContext(ctx).First(&stop, stopID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStopNotFound
		}
		return nil, err
	}
	stopResponse := mapStopToResponse(stop)
	return &stopResponse, nil
}

func (repo *StopRepository) ReSequenceStops(ctx context.Context, routeID uint, deletedStopSequence int) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		// Decrement the sequence numbers of all stops that had a higher sequence number than the deleted stop.
		return tx.Model(&models.Stop{}).
			Where("route_id = ? AND sequence > ?", routeID, deletedStopSequence).
			UpdateColumn("sequence", gorm.Expr("sequence - ?", 1)).Error
	})
}

func (repo *StopRepository) DeleteAndReSequenceStops(ctx context.Context, routeID uint, stopID uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		var stop models.Stop

		// Find the stop to get its sequence for later re-sequencing
		if err := tx.WithContext(ctx).Where("id = ? AND route_id = ?", stopID, routeID).First(&stop).Error; err != nil {
			return err
		}

		// Delete the stop
		if err := tx.WithContext(ctx).Delete(&models.Stop{}, stopID).Error; err != nil {
			return err
		}

		// Decrement the sequence numbers of all stops that had a higher sequence number than the deleted stop
		if err := tx.WithContext(ctx).Model(&models.Stop{}).
			Where("route_id = ? AND sequence > ?", routeID, stop.Sequence).
			UpdateColumn("sequence", gorm.Expr("sequence - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetStopByID retrieves a stop by its ID.
func (repo *StopRepository) GetStopByID(ctx context.Context, routeID uint, stopID uint) (*dto.StopResponse, error) {
	var stop models.Stop
	if err := repo.db.WithContext(ctx).
		Where("id = ? AND route_id = ?", stopID, routeID).
		First(&stop).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStopNotFound
		}
		return nil, err
	}

	stopResponse := mapStopToResponse(stop)
	return &stopResponse, nil

}

// GetStopBySequenceAndRouteID retrieves a stop by its sequence number and route ID.
func (repo *StopRepository) GetStopBySequenceAndRouteID(ctx context.Context, sequence int, routeID uint) (*models.Stop, error) {
	var stop models.Stop
	if err := repo.db.WithContext(ctx).
		Where("sequence = ? AND route_id = ?", sequence, routeID).
		First(&stop).Error; err != nil {
		return nil, err
	}
	return &stop, nil
}

// mapStopToResponse maps a models.Stop to a dto.StopResponse.
func mapStopToResponse(stop models.Stop) dto.StopResponse {
	return dto.StopResponse{
		StopID:   stop.ID,
		Name:     stop.Name,
		Sequence: stop.Sequence,
	}
}
