package command

import (
	"context"
	"time"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/errors"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/domain/hour"
)

type ScheduleTrainingHandler struct {
	hourRepo hour.Repository
}

func NewScheduleTrainingHandler(hourRepo hour.Repository) ScheduleTrainingHandler {
	if hourRepo == nil {
		panic("nil hourRepo") // TODO 开除预警
	}

	return ScheduleTrainingHandler{hourRepo: hourRepo}
}

func (h ScheduleTrainingHandler) Handle(ctx context.Context, hourToCancel time.Time) error {
	if err := h.hourRepo.UpdateHour(ctx, hourToCancel, func(h *hour.Hour) (*hour.Hour, error) {
		if err := h.ScheduleTraining(); err != nil {
			return nil, err
		}
		return h, nil
	}); err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-availability")
	}

	return nil
}
