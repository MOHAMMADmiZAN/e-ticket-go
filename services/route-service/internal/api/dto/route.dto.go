package dto

import (
	"route-service/internal/models"
	"time"
)

type RouteCreateRequest struct {
	Name          string    `json:"name" binding:"required"`
	StartTime     time.Time `json:"startTime" binding:"required"`
	Duration      int       `json:"duration" binding:"required,gt=0"` // Ensure duration is greater than 0
	StartLocation string    `json:"startLocation" binding:"required"`
	EndLocation   string    `json:"endLocation" binding:"required"`
}

type RouteInfo struct {
	RouteID         uint      `json:"route_id"`
	Name            string    `json:"name"`
	StartTime       time.Time `json:"start_time"`
	DurationMinutes int       `json:"duration_minutes"`
	StartLocation   string    `json:"start_location"`
	EndLocation     string    `json:"end_location"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type RouteResponse struct {
	ID            uint           `json:"id"`
	Name          string         `json:"name"`
	StartTime     time.Time      `json:"startTime"`
	Duration      int            `json:"duration"`
	StartLocation string         `json:"startLocation"`
	EndLocation   string         `json:"endLocation"`
	Stops         []StopResponse `json:"stops"` // Nested Stops within RouteResponse
	CreatedAt     string         `json:"createdAt"`
	UpdatedAt     string         `json:"updatedAt"`
}

// RouteUpdateRequest represents the request to update a route.
type RouteUpdateRequest struct {
	Name          *string    `json:"name,omitempty"`
	StartTime     *time.Time `json:"startTime,omitempty"`
	Duration      *int       `json:"duration,omitempty"`
	StartLocation *string    `json:"startLocation,omitempty"`
	EndLocation   *string    `json:"endLocation,omitempty"`
}

// ToModel write a function to covert dto to models
func (r *RouteCreateRequest) ToModel() *models.Route {
	return &models.Route{
		Name:          r.Name,
		StartTime:     r.StartTime,
		Duration:      r.Duration,
		StartLocation: r.StartLocation,
		EndLocation:   r.EndLocation,
	}
}

func (r *RouteUpdateRequest) ToModel(existingRouteResponse *RouteResponse) *models.Route {
	// Initialize a Route models with the values from the existing RouteResponse.
	updatedModel := &models.Route{
		Name:          existingRouteResponse.Name,
		StartTime:     existingRouteResponse.StartTime,
		Duration:      existingRouteResponse.Duration,
		StartLocation: existingRouteResponse.StartLocation,
		EndLocation:   existingRouteResponse.EndLocation,
		ID:            existingRouteResponse.ID,
	}

	// Overwrite the models fields with the values from the RouteUpdateRequest if provided.
	if r.Name != nil {
		updatedModel.Name = *r.Name
	}
	if r.StartTime != nil {
		updatedModel.StartTime = *r.StartTime
	}
	if r.Duration != nil {
		updatedModel.Duration = *r.Duration
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
		RouteID:         route.ID,
		Name:            route.Name,
		StartTime:       route.StartTime,
		DurationMinutes: route.Duration,
		StartLocation:   route.StartLocation,
		EndLocation:     route.EndLocation,
		CreatedAt:       route.CreatedAt,
		UpdatedAt:       route.UpdatedAt,
	}
}
