package service

import (
	"context"
	"e-ticket/services/route-service/internal/api/dto"
	"e-ticket/services/route-service/internal/model"
	"e-ticket/services/route-service/internal/repository"
	"errors"
	"fmt"
)

// Define business-specific error types.
var (
	ErrInvalidRouteID       = errors.New("the route ID provided is invalid for the specified stop")
	ErrStopNotFound         = errors.New("the stop was not found")
	ErrStopNameConflict     = errors.New("another stop with the same name exists on this route")
	ErrSequenceNotAvailable = errors.New("the sequence is not available for the stop")
)

// StopServiceInterface defines the service operations for stops.
type StopServiceInterface interface {
	AddStopToRoute(ctx context.Context, routeID uint, stop model.Stop) (*dto.StopResponse, error)
	GetStopsByRouteID(ctx context.Context, routeID uint) ([]dto.StopResponse, error)
	UpdateStop(ctx context.Context, routeID uint, stopID uint, stop model.Stop) (*dto.StopResponse, error)
	DeleteStop(ctx context.Context, routeID uint, stopID uint) error
}

// Ensure StopService adheres to the StopServiceInterface.
var _ StopServiceInterface = (*StopService)(nil)

// StopService provides operations for managing stops.
type StopService struct {
	stopRepo repository.StopRepositoryInterface
}

func NewStopService(stopRepo repository.StopRepositoryInterface) *StopService {
	return &StopService{
		stopRepo: stopRepo,
	}
}

// AddStopToRoute handles the business logic for adding a stop to a route.
func (s *StopService) AddStopToRoute(ctx context.Context, routeID uint, stop model.Stop) (*dto.StopResponse, error) {
	if stop.RouteID != routeID {
		return nil, ErrInvalidRouteID
	}

	// Check for stop name conflict within the route
	stops, err := s.GetStopsByRouteID(ctx, routeID)
	if err != nil {
		return nil, err
	}
	for _, existingStop := range stops {
		if existingStop.Name == stop.Name {
			return nil, ErrStopNameConflict
		}
		if existingStop.Sequence == stop.Sequence {
			return nil, ErrSequenceNotAvailable
		}
	}

	// Add the stop to the route
	return s.stopRepo.AddStopToRoute(ctx, stop)
}

// GetStopsByRouteID retrieves all stops associated with a route.
func (s *StopService) GetStopsByRouteID(ctx context.Context, routeID uint) ([]dto.StopResponse, error) {
	return s.stopRepo.ListAllStopsForRoute(ctx, routeID)
}
func newRouteIDMismatchError(expected, found uint) error {
	return fmt.Errorf("route ID mismatch: expected %v, found %v", expected, found)
}

// UpdateStop updates the details of a specific stop.
func (s *StopService) UpdateStop(ctx context.Context, routeID uint, stopID uint, updatedStop model.Stop) (*dto.StopResponse, error) {
	if updatedStop.RouteID != routeID {
		return nil, ErrInvalidRouteID
	}

	// Ensure the stop exists and belongs to the route
	existingStop, err := s.stopRepo.FindStopByID(ctx, stopID)
	if err != nil {
		return nil, ErrStopNotFound
	}
	if existingStop.Route.RouteID != routeID {
		return nil, ErrInvalidRouteID
	}

	// Ensure the sequence is available
	if existingStop.Sequence != updatedStop.Sequence {
		result, _ := s.stopRepo.GetStopBySequenceAndRouteID(ctx, updatedStop.Sequence, routeID)
		if result != nil {
			return nil, ErrSequenceNotAvailable
		}
	}

	// Check for name conflict
	stops, err := s.GetStopsByRouteID(ctx, routeID)
	if err != nil {
		return nil, err
	}
	for _, s := range stops {
		if s.StopID != stopID && s.Name == updatedStop.Name {
			return nil, ErrStopNameConflict
		}
	}
	// Update the stop details
	return s.stopRepo.UpdateStopDetails(ctx, updatedStop)
}

// DeleteStop removes a stop from a route.
func (s *StopService) DeleteStop(ctx context.Context, routeID uint, stopID uint) error {
	// The method DeleteAndReSequenceStops now encapsulates the whole transaction.
	return s.stopRepo.DeleteAndReSequenceStops(ctx, routeID, stopID)
}
