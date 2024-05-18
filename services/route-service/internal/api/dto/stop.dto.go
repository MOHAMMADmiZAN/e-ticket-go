package dto

import "time"

type AddStopRequest struct {
	Name     string `json:"name" binding:"required"`
	Sequence int    `json:"sequence" binding:"required,gt=0"`
}

type UpdateStopRequest struct {
	Name     string `json:"name" binding:"omitempty"`
	Sequence int    `json:"sequence" binding:"omitempty,gt=0"`
}

type StopResponse struct {
	StopID    uint               `json:"stop_id"`
	Name      string             `json:"name"`
	Sequence  int                `json:"sequence"`
	Schedules []ScheduleResponse `json:"schedules"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
