package repository

import (
	"context"
	"gorm.io/gorm"
	"route-service/internal/api/dto"
	"route-service/internal/model"
)

type ScheduleRepository interface {
	CreateSchedule(ctx context.Context, schedule *model.Schedule) (*dto.ScheduleResponse, error)
	GetScheduleByID(ctx context.Context, scheduleID uint) (*dto.ScheduleResponse, error)
	GetSchedulesByRouteID(ctx context.Context, routeID uint) ([]dto.ScheduleResponse, error)
	UpdateSchedule(ctx context.Context, schedule *model.Schedule) (*dto.ScheduleResponse, error)
	DeleteSchedule(ctx context.Context, scheduleID uint) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) CreateSchedule(ctx context.Context, schedule *model.Schedule) (*dto.ScheduleResponse, error) {
	if result := r.db.WithContext(ctx).Create(schedule); result.Error != nil {
		return nil, result.Error
	}
	return r.toScheduleResponse(schedule), nil
}

func (r *scheduleRepository) GetScheduleByID(ctx context.Context, scheduleID uint) (*dto.ScheduleResponse, error) {
	var schedule model.Schedule
	if result := r.db.WithContext(ctx).Preload("Stop").First(&schedule, scheduleID); result.Error != nil {
		return nil, result.Error
	}
	return r.toScheduleResponse(&schedule), nil
}

func (r *scheduleRepository) GetSchedulesByRouteID(ctx context.Context, routeID uint) ([]dto.ScheduleResponse, error) {
	var schedules []model.Schedule
	if result := r.db.WithContext(ctx).Where("route_id = ?", routeID).Find(&schedules); result.Error != nil {
		return nil, result.Error
	}

	scheduleResponses := make([]dto.ScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		scheduleResponses[i] = *r.toScheduleResponse(&schedule)
	}
	return scheduleResponses, nil
}

func (r *scheduleRepository) UpdateSchedule(ctx context.Context, schedule *model.Schedule) (*dto.ScheduleResponse, error) {
	if result := r.db.WithContext(ctx).Save(schedule); result.Error != nil {
		return nil, result.Error
	}
	return r.toScheduleResponse(schedule), nil
}

func (r *scheduleRepository) DeleteSchedule(ctx context.Context, scheduleID uint) error {
	return r.db.WithContext(ctx).Delete(&model.Schedule{}, scheduleID).Error
}

func (r *scheduleRepository) toScheduleResponse(schedule *model.Schedule) *dto.ScheduleResponse {
	if schedule == nil {
		return nil
	}
	if schedule.Stop.ID == 0 || schedule.Route.ID == 0 {
		// Load the Stop and Route explicitly if not preloaded
		err := r.db.Model(schedule).Association("Stop").Find(&schedule.Stop)
		if err != nil {
			return nil
		}
		err = r.db.Model(schedule).Association("Route").Find(&schedule.Route)
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

func (r *scheduleRepository) toRouteInfoDTO(route *model.Route) dto.RouteInfo {
	// Convert the model.Route to dto.RouteInfo. Please add the necessary fields.
	return dto.RouteInfo{
		RouteID:   route.ID,
		Name:      route.Name,
		CreatedAt: route.CreatedAt,
		UpdatedAt: route.UpdatedAt,
	}
}

func (r *scheduleRepository) toStopResponseDTO(stop *model.Stop) dto.StopResponse {
	// Convert the model.Stop to dto.StopResponse. Please add the necessary fields.
	return dto.StopResponse{
		StopID:    stop.ID,
		Name:      stop.Name,
		Sequence:  stop.Sequence,
		CreatedAt: stop.CreatedAt,
		UpdatedAt: stop.UpdatedAt,
	}
}
