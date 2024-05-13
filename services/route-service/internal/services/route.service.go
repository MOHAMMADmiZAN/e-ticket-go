package services

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"route-service/internal/api/dto"
	"route-service/internal/models"
	"route-service/internal/repository"
	"route-service/pkg"
)

// RouteService provides methods to work with the routes' repository.
type RouteService struct {
	repo        *repository.RouteRepository
	restyClient *resty.Client
}

// NewRouteService creates a new instance of RouteService.
func NewRouteService(repo *repository.RouteRepository) *RouteService {
	return &RouteService{
		repo:        repo,
		restyClient: resty.New(),
	}
}

// CreateRoute handles the creation of a new route.
func (s *RouteService) CreateRoute(ctx context.Context, route *models.Route) (*models.Route, error) {
	err := s.repo.Create(ctx, route)
	// Return  Route Response

	return route, err
}

// GetRoutes retrieves all routes from the database.
func (s *RouteService) GetRoutes(ctx context.Context) ([]dto.RouteResponse, error) {
	// Fetch all routes from the repository
	routes, err := s.repo.GetAll(ctx)
	if err != nil {
		// Return the error to the caller to handle (e.g., logging, retrying, etc.)
		return nil, err
	}

	// Preallocate a slice for the route responses with the same length as the routes slice
	routesResponse := make([]dto.RouteResponse, 0, len(routes))

	// Iterate over the fetched routes, mapping each to a dto.RouteResponse
	for _, route := range routes {
		// Note: Assuming MapRouteModelToRouteResponse accepts a pointer and returns a value or a pointer.
		// If it returns a value, this code works as is; if it returns a pointer, dereference it if needed.
		mappedRoute := MapRouteModelToRouteResponse(&route)
		// Append the mapped route response to the slice
		// If MapRouteModelToRouteResponse returns a pointer, use *mappedRoute; otherwise, use mappedRoute directly.
		routesResponse = append(routesResponse, *mappedRoute)
	}

	// Return the populated slice of route responses
	return routesResponse, nil
}

// GetRouteByID fetches a single route by its ID.
func (s *RouteService) GetRouteByID(ctx context.Context, id uint) (*dto.RouteResponse, error) {
	route, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return MapRouteModelToRouteResponse(route), nil
}

// UpdateRoute updates an existing route's details based on the provided request.
func (s *RouteService) UpdateRoute(ctx context.Context, route *models.Route) (*models.Route, error) {

	if err := s.repo.Update(ctx, route); err != nil {
		return nil, fmt.Errorf("error updating route: %v", err)
	}

	return route, nil
}

// DeleteRoute deletes a route by its ID.
func (s *RouteService) DeleteRoute(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func MapRouteModelToRouteResponse(route *models.Route) *dto.RouteResponse {
	stopsResponse := make([]dto.StopResponse, 0, len(route.Stops))
	for _, stop := range route.Stops {
		// Filter schedules for this specific stop
		stopsResponse = append(stopsResponse, dto.StopResponse{
			StopID:    stop.ID,
			Name:      stop.Name,
			Sequence:  stop.Sequence,
			CreatedAt: stop.CreatedAt,
			UpdatedAt: stop.UpdatedAt,
		})
	}

	return &dto.RouteResponse{
		ID:            route.ID,
		Name:          route.Name,
		StartTime:     route.StartTime,
		Duration:      route.Duration,
		StartLocation: route.StartLocation,
		EndLocation:   route.EndLocation,
		Stops:         stopsResponse,
		CreatedAt:     pkg.ConvertTime(route.CreatedAt),
		UpdatedAt:     pkg.ConvertTime(route.UpdatedAt),
	}
}
