package service

import (
	"context"
	"e-ticket/services/route-service/internal/api/dto"
	"e-ticket/services/route-service/internal/model"
	"e-ticket/services/route-service/internal/repository"
)

// RouteService provides methods to work with the routes repository.
type RouteService struct {
	repo *repository.RouteRepository
}

// NewRouteService creates a new instance of RouteService.
func NewRouteService(repo *repository.RouteRepository) *RouteService {
	return &RouteService{
		repo: repo,
	}
}

// CreateRoute handles the creation of a new route.
func (s *RouteService) CreateRoute(ctx context.Context, request dto.RouteCreateRequest) (*model.Route, error) {

	route := &model.Route{
		Name:          request.Name,
		StartTime:     request.StartTime,
		Duration:      request.Duration,
		StartLocation: request.StartLocation,
		EndLocation:   request.EndLocation,
	}

	err := s.repo.Create(ctx, route)
	return route, err
}

// GetRoutes retrieves all routes from the database.
func (s *RouteService) GetRoutes(ctx context.Context) ([]model.Route, error) {
	return s.repo.GetAll(ctx)
}

// GetRouteByID fetches a single route by its ID.
func (s *RouteService) GetRouteByID(ctx context.Context, id uint) (*model.Route, error) {
	return s.repo.GetByID(ctx, id)
}

// UpdateRoute updates an existing route's details.
//func (s *RouteService) UpdateRoute(ctx context.Context, id uint, name string, startTime time.Time, duration int, startLocation, endLocation string) (*model.Route, error) {
//	route := &model.Route{
//		Name:          name,
//		StartTime:     startTime,
//		Duration:      duration,
//		StartLocation: startLocation,
//		EndLocation:   endLocation,
//	}
//	err := s.repo.Update(ctx, id, route)
//	return route, err
//}

// DeleteRoute deletes a route by its ID.
func (s *RouteService) DeleteRoute(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
