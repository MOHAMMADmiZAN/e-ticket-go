package dto

import (
	"errors"
	"route-service/internal/models"
	"time"
)

type ScheduleResponse struct {
	ScheduleID    uint      `json:"schedule_id"`
	ArrivalTime   time.Time `json:"arrival_time"`
	DepartureTime time.Time `json:"departure_time"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AddScheduleRequest struct {
	StopID        uint      `json:"stop_id" binding:"required"`
	ArrivalTime   time.Time `json:"arrival_time" binding:"required"`
	DepartureTime time.Time `json:"departure_time" binding:"required"`
}

// Validate checks the validity of the AddScheduleRequest fields.
func (a *AddScheduleRequest) Validate() error {
	if a.StopID == 0 {
		return errors.New("stop ID is required and must be greater than zero")
	}
	if a.ArrivalTime.IsZero() {
		return errors.New("arrival time is required and must be a valid date and time")
	}
	if a.DepartureTime.IsZero() {
		return errors.New("departure time is required and must be a valid date and time")
	}
	if a.ArrivalTime.After(a.DepartureTime) {
		return errors.New("arrival time must be earlier than departure time")
	}
	return nil // No error means the request is valid.
}

// ToModel converts AddScheduleRequest to the Schedule models.
func (a *AddScheduleRequest) ToModel() *models.Schedule {
	return &models.Schedule{
		StopID:        a.StopID,
		ArrivalTime:   a.ArrivalTime,
		DepartureTime: a.DepartureTime,
	}
}

type UpdateScheduleRequest struct {
	StopID        uint      `json:"stop_id" binding:"required"`
	ArrivalTime   time.Time `json:"arrival_time" binding:"required"`
	DepartureTime time.Time `json:"departure_time" binding:"required"`
}

// Validate checks the validity of the UpdateScheduleRequest fields.
func (u *UpdateScheduleRequest) Validate() error {
	if u.StopID == 0 {
		return errors.New("stop ID is required and must be greater than zero")
	}
	if u.ArrivalTime.IsZero() || u.DepartureTime.IsZero() {
		return errors.New("both arrival time and departure time must be provided and be valid dates and times")
	}
	if u.ArrivalTime.After(u.DepartureTime) {
		return errors.New("arrival time must be earlier than departure time")
	}

	return nil // No error means the request is valid.
}

// ToModel converts UpdateScheduleRequest to the Schedule models.
func (u *UpdateScheduleRequest) ToModel() *models.Schedule {
	return &models.Schedule{
		StopID:        u.StopID,
		ArrivalTime:   u.ArrivalTime,
		DepartureTime: u.DepartureTime,
	}
}
