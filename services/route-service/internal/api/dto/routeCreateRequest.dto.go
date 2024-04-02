package dto

import "time"

type RouteCreateRequest struct {
	Name          string    `json:"name" binding:"required"`
	StartTime     time.Time `json:"startTime" binding:"required"`
	Duration      int       `json:"duration" binding:"required,gt=0"` // Ensure duration is greater than 0
	StartLocation string    `json:"startLocation" binding:"required"`
	EndLocation   string    `json:"endLocation" binding:"required"`
}
