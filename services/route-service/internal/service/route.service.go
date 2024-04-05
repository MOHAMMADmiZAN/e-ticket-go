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

func MapRouteModelToRouteResponse(route *model.Route) *dto.RouteResponse {
	stopsResponse := make([]dto.StopResponse, 0, len(route.Stops))

	// Assuming Schedules are directly related to the Route and not nested under Stops in your model.
	// We'll need to filter these Schedules to match them with their respective Stops.
	for _, stop := range route.Stops {
		// Filter schedules for this specific stop
		filteredSchedules := filterSchedulesForStop(route.Schedules, stop.ID)

		schedulesResponse := make([]dto.ScheduleResponse, 0, len(filteredSchedules))
		for _, schedule := range filteredSchedules {
			schedulesResponse = append(schedulesResponse, dto.ScheduleResponse{
				ScheduleID:    schedule.ID,
				ArrivalTime:   schedule.ArrivalTime,
				DepartureTime: schedule.DepartureTime,
				CreatedAt:     schedule.CreatedAt,
				UpdatedAt:     schedule.UpdatedAt,
			})
		}

		stopsResponse = append(stopsResponse, dto.StopResponse{
			StopID:    stop.ID,
			Name:      stop.Name,
			Sequence:  stop.Sequence,
			Schedules: schedulesResponse,
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
		CreatedAt:     route.CreatedAt,
		UpdatedAt:     route.UpdatedAt,
	}
}

// filterSchedulesForStop takes a slice of all Schedules related to the Route and a StopID,
// and returns a filtered slice of Schedules that belong to the Stop.
func filterSchedulesForStop(allSchedules []model.Schedule, stopID uint) []model.Schedule {
	filtered := make([]model.Schedule, 0)
	for _, schedule := range allSchedules {
		if schedule.StopID == stopID {
			filtered = append(filtered, schedule)
		}
	}
	return filtered
}
