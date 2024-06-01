package repository

import (
	"context"
	"route-service/internal/api/dto"
	"route-service/internal/models"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(ctx context.Context, schedule *models.Schedule) (*dto.ScheduleResponse, error)
	GetScheduleByID(ctx context.Context, scheduleID uint, stopId uint) (*dto.ScheduleResponse, error)
	GetSchedules(ctx context.Context, stopId uint) ([]dto.ScheduleResponse, error)
	UpdateSchedule(ctx context.Context, schedule *models.Schedule) (*dto.ScheduleResponse, error)
	DeleteSchedule(ctx context.Context, scheduleID uint, stopId uint) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) CreateSchedule(ctx context.Context, schedule *models.Schedule) (*dto.ScheduleResponse, error) {
	if result := r.db.WithContext(ctx).Create(schedule); result.Error != nil {
		return nil, result.Error
	}
	return r.toScheduleResponse(schedule), nil
}

func (r *scheduleRepository) GetScheduleByID(ctx context.Context, scheduleID uint, stopId uint) (*dto.ScheduleResponse, error) {
	var schedule models.Schedule
	if result := r.db.WithContext(ctx).Where("stop_id = ?", stopId).First(&schedule, scheduleID); result.Error != nil {
		return nil, result.Error
	}
	return r.toScheduleResponse(&schedule), nil
}

func (r *scheduleRepository) GetSchedules(ctx context.Context, stopId uint) ([]dto.ScheduleResponse, error) {
	var schedules []models.Schedule
	if result := r.db.WithContext(ctx).Where("stop_id = ?", stopId).Find(&schedules); result.Error != nil {
		return nil, result.Error
	}

	scheduleResponses := make([]dto.ScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		scheduleResponses[i] = *r.toScheduleResponse(&schedule)
	}
	return scheduleResponses, nil
}

func (r *scheduleRepository) UpdateSchedule(ctx context.Context, schedule *models.Schedule) (*dto.ScheduleResponse, error) {
	if result := r.db.WithContext(ctx).Save(schedule); result.Error != nil {
		return nil, result.Error
	}
	return r.toScheduleResponse(schedule), nil
}

// DeleteSchedule deletes a schedule by its ID and route ID.
func (r *scheduleRepository) DeleteSchedule(ctx context.Context, scheduleID uint, stopId uint) error {
	return r.db.WithContext(ctx).Where("stop_id = ?", stopId).Delete(&models.Schedule{}, scheduleID).Error
}

func (r *scheduleRepository) toScheduleResponse(schedule *models.Schedule) *dto.ScheduleResponse {
	if schedule == nil {
		return nil
	}
	if schedule.Stop.ID == 0 {
		// Load the Stop and Route explicitly if not preloaded
		err := r.db.Model(schedule).Association("Stop").Find(&schedule.Stop)
		if err != nil {
			return nil
		}
	}
	return &dto.ScheduleResponse{
		ScheduleID:    schedule.ID,
		ArrivalTime:   schedule.ArrivalTime,
		DepartureTime: schedule.DepartureTime,
		CreatedAt:     schedule.CreatedAt,
		UpdatedAt:     schedule.UpdatedAt,
	}
}

func (r *scheduleRepository) toStopResponseDTO(stop *models.Stop) dto.StopResponse {
	// Convert the models.Stop to dto.StopResponse. Please add the necessary fields.
	return dto.StopResponse{
		StopID:   stop.ID,
		Name:     stop.Name,
		Sequence: stop.Sequence,
	}
}
