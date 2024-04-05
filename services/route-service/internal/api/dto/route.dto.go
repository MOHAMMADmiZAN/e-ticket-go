package dto

import "time"

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
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}
