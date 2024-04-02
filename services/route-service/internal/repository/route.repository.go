package repository

import (
	"context"
	"e-ticket/services/route-service/internal/model"
	"gorm.io/gorm"
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
	err := r.db.WithContext(ctx).Find(&routes).Error
	return routes, err
}

// GetByID fetches a single route record by its ID from the database.
func (r *RouteRepository) GetByID(ctx context.Context, id uint) (*model.Route, error) {
	var route model.Route
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&route).Error
	if err != nil {
		return nil, err
	}
	return &route, nil
}

// Update updates an existing route record in the database.
func (r *RouteRepository) Update(ctx context.Context, route *model.Route) error {
	return r.db.WithContext(ctx).Save(route).Error
}

// Delete removes a route record from the database.
func (r *RouteRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Route{}, id).Error
}
