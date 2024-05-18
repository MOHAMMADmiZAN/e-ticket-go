package services

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"route-service/internal/api/dto"
	"route-service/internal/models"
	"route-service/internal/repository"

	"github.com/go-resty/resty/v2"
)

// RouteService provides methods to work with the routes' repository.
type RouteService struct {
	repo        *repository.RouteRepository
	restyClient *resty.Client
}

var (
	busServiceBaseURL = os.Getenv("BUS_SERVICE_BASE_URL")
)

type BusServiceResponse struct {
	Success bool              `json:"success"`
	Data    []dto.BusResponse `json:"data"`
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
		mappedRoute := MapRouteModelToRouteResponse(&route, s, false)
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
	return MapRouteModelToRouteResponse(route, s, true), nil
}

// UpdateRoute updates an existing route's details based on the provided request.
func (s *RouteService) UpdateRoute(ctx context.Context, route *models.Route) (*dto.RouteResponse, error) {

	if err := s.repo.Update(ctx, route); err != nil {
		return nil, fmt.Errorf("error updating route: %v", err)
	}

	return MapRouteModelToRouteResponse(route, s), nil
}

// DeleteRoute deletes a route by its ID.
func (s *RouteService) DeleteRoute(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// MapRouteModelToRouteResponse maps the route model to the route response DTO
func MapRouteModelToRouteResponse(route *models.Route, s *RouteService, isExPend ...bool) *dto.RouteResponse {
	expand := false
	if len(isExPend) > 0 {
		expand = isExPend[0]
	}

	routeResponse := &dto.RouteResponse{
		ID:            route.ID,
		Name:          route.Name,
		StartTime:     route.StartTime,
		Duration:      route.Duration,
		StartLocation: route.StartLocation,
		EndLocation:   route.EndLocation,
		// CreatedAt:     pkg.ConvertTime(route.CreatedAt),
		// UpdatedAt:     pkg.ConvertTime(route.UpdatedAt),
	}

	if expand {
		// Process stops
		stopsResponse := make([]dto.StopResponse, len(route.Stops))
		for i, stop := range route.Stops {
			stopsResponse[i] = dto.StopResponse{
				Name:     stop.Name,
				Sequence: stop.Sequence,
			}
		}
		routeResponse.Stops = stopsResponse

		// Get buses for this route
		buses, err := getBuses(s, route.ID)
		if err != nil {
			// Log the error and continue with an empty bus list
			fmt.Printf("error getting buses for route %d: %v\n", route.ID, err)
			buses = []dto.BusResponse{}
		}
		routeResponse.Buses = buses
	}

	return routeResponse
}

func getBuses(s *RouteService, routeID uint) ([]dto.BusResponse, error) {
	var busServiceResponse BusServiceResponse
	url := fmt.Sprintf("%s/routes/%d", busServiceBaseURL, routeID)
	resp, err := s.restyClient.R().
		SetResult(&busServiceResponse).
		Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("bus Service responded with status code: %d", resp.StatusCode())
	}

	if !busServiceResponse.Success {
		return nil, fmt.Errorf("bus Service responded with success: false")
	}

	return busServiceResponse.Data, nil
}
