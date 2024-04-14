package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"route-service/internal/model"
)

// RouteRepository is responsible for handling the operations related to the Route model.
type RouteRepository struct {
	db *gorm.DB
}

// NewRouteRepository creates a new instance of RouteRepository.
func NewRouteRepository(db *gorm.DB) *RouteRepository {
	return &RouteRepository{
		db: db,
	}
}

// Create creates a new route record in the database.
func (r *RouteRepository) Create(ctx context.Context, route *model.Route) error {
	return r.db.WithContext(ctx).Create(route).Error
}

// GetAll fetches all route records from the database.
func (r *RouteRepository) GetAll(ctx context.Context) ([]model.Route, error) {
	var routes []model.Route
	err := r.db.WithContext(ctx).
		Preload("Stops", func(db *gorm.DB) *gorm.DB {
			return db.Order("stops.sequence ASC") // Order stops by sequence
		}).
		Preload("Schedules", func(db *gorm.DB) *gorm.DB {
			return db.Order("schedules.departure_time ASC") // Order schedules by departure time
		}).
		Find(&routes).Error
	return routes, err
}

// GetByID fetches a single route record by its ID from the database.
func (r *RouteRepository) GetByID(ctx context.Context, id uint) (*model.Route, error) {
	var route model.Route
	err := r.db.WithContext(ctx).
		Preload("Stops", func(db *gorm.DB) *gorm.DB {
			return db.Order("stops.sequence ASC") // Order stops by sequence
		}).
		Preload("Schedules", func(db *gorm.DB) *gorm.DB {
			return db.Order("schedules.departure_time ASC") // Order schedules by departure time
		}).
		Where("id = ?", id).
		First(&route).Error
	if err != nil {
		return nil, err
	}
	return &route, nil
}

// Update updates an existing route record in the database.
func (r *RouteRepository) Update(ctx context.Context, route *model.Route) error {
	if err := r.db.WithContext(ctx).Model(&model.Route{}).Where("id = ?", route.ID).Updates(route).Error; err != nil {
		// Logging the error with context (like request ID if available) would be beneficial for debugging
		log.Printf("Failed to update route with ID %d: %v", route.ID, err)
		return fmt.Errorf("update failed: %w", err)
	}
	return nil
}

func (r *RouteRepository) Delete(ctx context.Context, id uint) error {
	// Start a new transaction
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Retrieve and delete the route using the primary key, `id`.
		if err := tx.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&model.Route{}).Error; err != nil {
			// Returning any error will rollback the transaction
			return err
		}
		return nil
	})

	if err != nil {
		log.Printf("Failed to delete route with ID %d: %v", id, err)
		return err
	}

	return nil
}
