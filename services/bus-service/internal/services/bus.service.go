package services

import (
	"bus-service/internal/api/dto"
	"bus-service/internal/models"
	"bus-service/internal/repository"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

var (
	bookingServiceBaseURL = "http://booking-service"
	routeServiceBaseURL   = os.Getenv("ROUTE_SERVICE_BASE_URL")
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

func (service *BusService) getRoute(routeID uint) (*dto.RouteResponse, error) {
	var route dto.RouteResponse
	url := fmt.Sprintf("%s/%d", routeServiceBaseURL, routeID)
	resp, err := service.restyClient.R().
		SetResult(&route).
		Get(url)

	if err != nil {
		return nil, err // Return error if the API call fails
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("RouteService responded with status code: %d", resp.StatusCode())
	}

	return &route, nil // Return the route if everything is okay
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
	_, err := service.getRoute(bus.RouteID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify route status: %v", err.Error())
	}
	// Check bus with the same LicensePlate
	busWithLicensePlate, err := service.busRepo.GetBusByCodeOrLicensePlate(bus.BusCode, bus.LicensePlate)
	if err != nil {
		// check if the error is not found goORM NotFound error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to check for existing bus: %v", err)
		} else {
			log.Print(`err`, err)
		}
	}
	if busWithLicensePlate != nil {
		return nil, errors.New("bus with the same license plate Or same bus code already exists")
	}
	return service.busRepo.CreateBus(bus)
}

// UpdateBus updates the details of an existing bus.
func (service *BusService) UpdateBus(bus models.Bus) (*models.Bus, error) {
	_, err := service.getRoute(bus.RouteID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify route status: %v", err)
	}

	return service.busRepo.UpdateBus(bus)
}

// DeleteBus removes a bus from the system after checking for future bookings.
func (service *BusService) DeleteBus(busID uint) error {
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

// GetBusesByRoute retrieves all buses by their route ID.
func (service *BusService) GetBusesByRoute(routeID uint) ([]models.Bus, error) {
	return service.busRepo.GetBusesByRouteID(routeID)
}
