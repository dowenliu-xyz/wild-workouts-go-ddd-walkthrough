package command

import (
	"context"
	"time"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/errors"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/domain/hour"
)

type MakeHoursUnavailableHandler struct {
	hourRepo hour.Repository
}

func NewMakeHoursUnavailableHandler(hourRepo hour.Repository) MakeHoursUnavailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil") // TODO 开除预警
	}

	return MakeHoursUnavailableHandler{hourRepo: hourRepo}
}

func (c MakeHoursUnavailableHandler) Handle(ctx context.Context, hours []time.Time) error {
	for _, hourToUpdate := range hours {
		if err := c.hourRepo.UpdateHour(ctx, hourToUpdate, func(h *hour.Hour) (*hour.Hour, error) {
			if err := h.MakeNotAvailable(); err != nil {
				return nil, err
			}
			return h, nil
		}); err != nil {
			return errors.NewSlugError(err.Error(), "unable-to-update-availability")
		}
	}

	return nil
}
