package dto

import "time"

type AddStopRequest struct {
	Name     string `json:"name" binding:"required"`
	Sequence int    `json:"sequence" binding:"required,gt=0"`
}

type UpdateStopRequest struct {
	Name     string `json:"name" binding:"required"`
	Sequence int    `json:"sequence" binding:"required,gt=0"`
}

type StopResponse struct {
	StopID    uint      `json:"stop_id"`
	Name      string    `json:"name"`
	Sequence  int       `json:"sequence"`
	Route     RouteInfo `json:"route"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
