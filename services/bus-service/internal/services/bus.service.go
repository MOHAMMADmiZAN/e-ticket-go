package services

import (
	"bus-service/internal/models"
	"bus-service/internal/repository"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

const (
	routeServiceBaseURL   = "http://route-service"
	bookingServiceBaseURL = "http://booking-service"
)

type BusService struct {
	busRepo     *repository.BusRepository
	restyClient *resty.Client
}

func NewBusService(busRepo *repository.BusRepository) *BusService {
	return &BusService{
		busRepo:     busRepo,
		restyClient: resty.New(),
	}
}

// RouteServiceResponse represents the response from the RouteService for an active check.
type RouteServiceResponse struct {
	Active bool `json:"active"`
}

// BookingServiceResponse represents the response from the BookingService for future bookings check.
type BookingServiceResponse struct {
	HasBookings bool `json:"hasBookings"`
}

func (service *BusService) isRouteActive(routeID uint) (bool, error) {
	var response RouteServiceResponse
	resp, err := service.restyClient.R().
		SetResult(&response).
		Get(fmt.Sprintf("%s/routes/%d/isActive", routeServiceBaseURL, routeID))

	if err != nil {
		return false, err
	}
	if resp.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("RouteService responded with status code: %d", resp.StatusCode())
	}
	return response.Active, nil
}

func (service *BusService) hasFutureBookings(busID uint) (bool, error) {
	var response BookingServiceResponse
	resp, err := service.restyClient.R().
		SetResult(&response).
		Get(fmt.Sprintf("%s/buses/%d/hasFutureBookings", bookingServiceBaseURL, busID))

	if err != nil {
		return false, err
	}
	if resp.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("BookingService responded with status code: %d", resp.StatusCode())
	}
	return response.HasBookings, nil
}

// GetAllBuses retrieves all buses.
func (service *BusService) GetAllBuses() ([]models.Bus, error) {
	buses, err := service.busRepo.GetAllBuses()
	if err != nil {
		return nil, err
	}
	return buses, nil
}

// GetBusByID retrieves a bus by its ID.
func (service *BusService) GetBusByID(id uint) (*models.Bus, error) {
	bus, err := service.busRepo.GetBusByID(id)
	if err != nil {
		return nil, err
	}
	return bus, nil
}

// CreateBus adds a new bus to the system after validating the route.
func (service *BusService) CreateBus(bus models.Bus) (*models.Bus, error) {
	active, err := service.isRouteActive(bus.RouteID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify route status: %v", err)
	}
	if !active {
		return nil, errors.New("cannot create bus on an inactive route")
	}
	return service.busRepo.CreateBus(bus)
}

// UpdateBus updates the details of an existing bus.
func (service *BusService) UpdateBus(bus models.Bus) (*models.Bus, error) {
	active, err := service.isRouteActive(bus.RouteID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify route status: %v", err)
	}
	if !active {
		return nil, errors.New("cannot update bus on an inactive route")
	}
	return service.busRepo.UpdateBus(bus)
}

// DeleteBus removes a bus from the system after checking for future bookings.
func (service *BusService) DeleteBus(busID uint) error {
	hasBookings, err := service.hasFutureBookings(busID)
	if err != nil {
		return fmt.Errorf("failed to check for future bookings: %v", err)
	}
	if hasBookings {
		return errors.New("cannot delete a bus with future bookings")
	}
	return service.busRepo.DeleteBus(busID)
}

// GetBusesByStatus retrieves buses by their operational status.
func (service *BusService) GetBusesByStatus(status string) ([]models.Bus, error) {
	buses, err := service.busRepo.GetBusesByStatus(status)
	if err != nil {
		return nil, err
	}
	return buses, nil
}

// UpdateBusServiceDates updates the service dates for a bus after validation.
func (service *BusService) UpdateBusServiceDates(id uint, lastServiceDate, nextServiceDate time.Time) error {
	if lastServiceDate.After(nextServiceDate) {
		return errors.New("last service date must be before the next service date")
	}
	return service.busRepo.UpdateBusServiceDates(id, lastServiceDate, nextServiceDate)
}
