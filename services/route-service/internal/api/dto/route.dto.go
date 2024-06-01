package dto

import (
	"route-service/internal/models"
	"time"
)

type RouteCreateRequest struct {
	Name          string `json:"name" binding:"required"`
	StartLocation string `json:"startLocation" binding:"required"`
	EndLocation   string `json:"endLocation" binding:"required"`
}

type RouteInfo struct {
	RouteID       uint      `json:"route_id"`
	Name          string    `json:"name"`
	StartLocation string    `json:"start_location"`
	EndLocation   string    `json:"end_location"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RouteResponse struct {
	ID            uint           `json:"id"`
	Name          string         `json:"name"`
	StartLocation string         `json:"startLocation"`
	EndLocation   string         `json:"endLocation"`
	Stops         []StopResponse `json:"stops,omitempty"` // Optional Stops
	Buses         []BusResponse  `json:"buses,omitempty"` // Optional Buses
	//CreatedAt string        `json:"createdAt"`
	//UpdatedAt string        `json:"updatedAt"`
}

// RouteUpdateRequest represents the request to update a route.
type RouteUpdateRequest struct {
	Name          *string `json:"name,omitempty"`
	StartLocation *string `json:"startLocation,omitempty"`
	EndLocation   *string `json:"endLocation,omitempty"`
}

// ToModel write a function to covert dto to models
func (r *RouteCreateRequest) ToModel() *models.Route {
	return &models.Route{
		Name:          r.Name,
		StartLocation: r.StartLocation,
		EndLocation:   r.EndLocation,
	}
}

func (r *RouteUpdateRequest) ToModel(existingRouteResponse *RouteResponse) *models.Route {
	// Initialize a Route models with the values from the existing RouteResponse.
	updatedModel := &models.Route{
		Name:          existingRouteResponse.Name,
		StartLocation: existingRouteResponse.StartLocation,
		EndLocation:   existingRouteResponse.EndLocation,
		ID:            existingRouteResponse.ID,
	}

	// Overwrite the models fields with the values from the RouteUpdateRequest if provided.
	if r.Name != nil {
		updatedModel.Name = *r.Name
	}
	if r.StartLocation != nil {
		updatedModel.StartLocation = *r.StartLocation
	}
	if r.EndLocation != nil {
		updatedModel.EndLocation = *r.EndLocation
	}

	return updatedModel
}

func RouteModelToRouteInfo(route *models.Route) RouteInfo {
	return RouteInfo{
		RouteID:       route.ID,
		Name:          route.Name,
		StartLocation: route.StartLocation,
		EndLocation:   route.EndLocation,
		CreatedAt:     route.CreatedAt,
		UpdatedAt:     route.UpdatedAt,
	}
}

type BusResponse struct {
	ID              uint      `json:"id"`
	BusCode         string    `json:"busCode"`
	Capacity        int       `json:"capacity"`
	LicensePlate    string    `json:"licensePlate"`
	Status          string    `json:"status"`
	LastServiceDate time.Time `json:"lastServiceDate"`
	NextServiceDate time.Time `json:"nextServiceDate"`
}
