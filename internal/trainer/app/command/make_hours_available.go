package command

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/decorator"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/errors"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/domain/hour"
)

type MakeHoursAvailable struct {
	Hours []time.Time
}

type MakeHoursAvailableHandler decorator.CommandHandler[MakeHoursAvailable]

type makeHoursAvailableHandler struct {
	hourRepo hour.Repository
}

func NewMakeHoursAvailableHandler(
	hourRepo hour.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) MakeHoursAvailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil") // TODO 开除预警
	}

	return decorator.ApplyCommandDecorators[MakeHoursAvailable](
		makeHoursAvailableHandler{hourRepo: hourRepo},
		logger,
		metricsClient,
	)
}

func (c makeHoursAvailableHandler) Handle(ctx context.Context, cmd MakeHoursAvailable) error {
	for _, hourToUpdate := range cmd.Hours {
		if err := c.hourRepo.UpdateHour(ctx, hourToUpdate, func(h *hour.Hour) (*hour.Hour, error) {
			if err := h.MakeAvailable(); err != nil {
				return nil, err
			}
			return h, nil
		}); err != nil {
			return errors.NewSlugError(err.Error(), "unable-to-update-availability")
		}
	}

	return nil
}
