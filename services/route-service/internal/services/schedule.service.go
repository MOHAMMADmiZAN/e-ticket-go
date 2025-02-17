package services

import (
	"context"
	"fmt"
	"route-service/internal/api/dto"
	"route-service/internal/models"
	"route-service/internal/repository"
	"time"
)

type ScheduleService interface {
	CreateSchedule(ctx context.Context, newSchedule models.Schedule) (*dto.ScheduleResponse, error)
	GetScheduleByID(ctx context.Context, scheduleID uint, stopId uint) (*dto.ScheduleResponse, error)
	GetSchedules(ctx context.Context, stopId uint) ([]dto.ScheduleResponse, error)
	UpdateSchedule(ctx context.Context, scheduleID uint, updateSchedule models.Schedule) (*dto.ScheduleResponse, error)
	DeleteSchedule(ctx context.Context, scheduleID uint, stopId uint) error
}

type scheduleService struct {
	repo     repository.ScheduleRepository
	cache    map[uint]*dto.ScheduleResponse
	cacheTTL time.Duration
}

func NewScheduleService(repo repository.ScheduleRepository, cacheTTL time.Duration) ScheduleService {
	return &scheduleService{
		repo:     repo,
		cache:    make(map[uint]*dto.ScheduleResponse),
		cacheTTL: cacheTTL,
	}
}

func (s *scheduleService) CreateSchedule(ctx context.Context, newSchedule models.Schedule) (*dto.ScheduleResponse, error) {

	resp, err := s.repo.CreateSchedule(ctx, &newSchedule)
	if err != nil {
		// Log repository error
		return nil, fmt.Errorf("creating schedule: %w", err)
	}

	// Cache the new schedule
	s.cache[resp.ScheduleID] = resp

	// Return the response
	return resp, nil
}

func (s *scheduleService) GetScheduleByID(ctx context.Context, scheduleID uint, stopId uint) (*dto.ScheduleResponse, error) {
	// Check if the schedule is in the cache
	if cachedSchedule, found := s.cache[scheduleID]; found {
		// Log cache hit
		return cachedSchedule, nil
	}
	// Log cache miss and fetch from repository
	resp, err := s.repo.GetScheduleByID(ctx, scheduleID, stopId)

	if err != nil {
		// Log repository error
		return nil, fmt.Errorf("getting schedule by ID: %w", err)
	}

	// Cache the fetched schedule
	s.cache[resp.ScheduleID] = resp

	// Return the schedule
	return resp, nil
}

func (s *scheduleService) GetSchedules(ctx context.Context, stopId uint) ([]dto.ScheduleResponse, error) {
	// This can also be cached if required
	resp, err := s.repo.GetSchedules(ctx, stopId)
	if err != nil {
		// Log repository error
		return nil, fmt.Errorf("getting schedules by route ID: %w", err)
	}

	// Return the schedules
	return resp, nil
}

func (s *scheduleService) UpdateSchedule(ctx context.Context, scheduleID uint, updateSchedule models.Schedule) (*dto.ScheduleResponse, error) {

	resp, err := s.repo.UpdateSchedule(ctx, &updateSchedule)
	if err != nil {
		// Log repository error
		return nil, fmt.Errorf("updating schedule: %w", err)
	}

	// Update the cache with the new data
	s.cache[resp.ScheduleID] = resp

	// Return the updated schedule
	return resp, nil
}

func (s *scheduleService) DeleteSchedule(ctx context.Context, scheduleID uint, stopId uint) error {
	err := s.repo.DeleteSchedule(ctx, scheduleID, stopId)
	if err != nil {
		// Log repository error
		return fmt.Errorf("deleting schedule: %w", err)
	}

	// Remove the schedule from the cache
	delete(s.cache, scheduleID)

	// Return success
	return nil
}
