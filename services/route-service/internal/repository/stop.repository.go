package repository

import (
	"context"
	"e-ticket/services/route-service/internal/api/dto"
	"e-ticket/services/route-service/internal/model"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrStopNotFound = errors.New("stop not found")
	// You can define more specific errors related to database operations if needed.
)

// StopRepositoryInterface defines the methods that our StopRepository must implement.
type StopRepositoryInterface interface {
	AddStopToRoute(ctx context.Context, stop model.Stop) (*dto.StopResponse, error)
	ListAllStopsForRoute(ctx context.Context, routeID uint) ([]dto.StopResponse, error)
	UpdateStopDetails(ctx context.Context, stop model.Stop) (*dto.StopResponse, error)
	DeleteStop(ctx context.Context, routeID uint, stopID uint) error
	FindStopByID(ctx context.Context, stopID uint) (*dto.StopResponse, error)
	ReSequenceStops(ctx context.Context, routeID uint, deletedStopSequence int) error
	DeleteAndReSequenceStops(ctx context.Context, routeID uint, stopID uint) error
	GetStopBySequenceAndRouteID(ctx context.Context, sequence int, routeID uint) (*model.Stop, error)
}

// Ensure that StopRepository implements the StopRepositoryInterface.
var _ StopRepositoryInterface = (*StopRepository)(nil)

type StopRepository struct {
	db *gorm.DB
}

func NewStopRepository(db *gorm.DB) *StopRepository {
	return &StopRepository{db: db}
}

func (repo *StopRepository) AddStopToRoute(ctx context.Context, stop model.Stop) (*dto.StopResponse, error) {
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
	var stops []model.Stop
	if err := repo.db.WithContext(ctx).
		Preload("Route").
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

func (repo *StopRepository) UpdateStopDetails(ctx context.Context, stop model.Stop) (*dto.StopResponse, error) {
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
	if err := repo.db.WithContext(ctx).Where("id = ? AND route_id = ?", stopID, routeID).Delete(&model.Stop{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *StopRepository) FindStopByID(ctx context.Context, stopID uint) (*dto.StopResponse, error) {
	var stop model.Stop
	if err := repo.db.WithContext(ctx).Preload("Route").First(&stop, stopID).Error; err != nil {
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
		return tx.Model(&model.Stop{}).
			Where("route_id = ? AND sequence > ?", routeID, deletedStopSequence).
			UpdateColumn("sequence", gorm.Expr("sequence - ?", 1)).Error
	})
}

func (repo *StopRepository) DeleteAndReSequenceStops(ctx context.Context, routeID uint, stopID uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		var stop model.Stop

		// Find the stop to get its sequence for later re-sequencing
		if err := tx.WithContext(ctx).Where("id = ? AND route_id = ?", stopID, routeID).First(&stop).Error; err != nil {
			return err
		}

		// Delete the stop
		if err := tx.WithContext(ctx).Delete(&model.Stop{}, stopID).Error; err != nil {
			return err
		}

		// Decrement the sequence numbers of all stops that had a higher sequence number than the deleted stop
		if err := tx.WithContext(ctx).Model(&model.Stop{}).
			Where("route_id = ? AND sequence > ?", routeID, stop.Sequence).
			UpdateColumn("sequence", gorm.Expr("sequence - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetStopBySequenceAndRouteID retrieves a stop by its sequence number and route ID.
func (repo *StopRepository) GetStopBySequenceAndRouteID(ctx context.Context, sequence int, routeID uint) (*model.Stop, error) {
	var stop model.Stop
	if err := repo.db.WithContext(ctx).
		Where("sequence = ? AND route_id = ?", sequence, routeID).
		First(&stop).Error; err != nil {
		return nil, err
	}
	return &stop, nil
}

// mapStopToResponse maps a model.Stop to a dto.StopResponse.
func mapStopToResponse(stop model.Stop) dto.StopResponse {
	return dto.StopResponse{
		StopID:    stop.ID,
		Name:      stop.Name,
		Sequence:  stop.Sequence,
		Route:     mapRouteToInfo(stop.Route),
		CreatedAt: stop.CreatedAt,
		UpdatedAt: stop.UpdatedAt,
	}
}

// mapRouteToInfo maps a model.Route to a dto.RouteInfo.
func mapRouteToInfo(route model.Route) dto.RouteInfo {
	return dto.RouteInfo{
		RouteID:         route.ID,
		Name:            route.Name,
		StartTime:       route.StartTime,
		DurationMinutes: route.Duration,
		StartLocation:   route.StartLocation,
		EndLocation:     route.EndLocation,
		CreatedAt:       route.CreatedAt,
		UpdatedAt:       route.UpdatedAt,
	}
}
