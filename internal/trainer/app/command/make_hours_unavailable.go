package command

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/decorator"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/errors"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/domain/hour"
)

type MakeHoursUnavailable struct {
	Hours []time.Time
}

type MakeHoursUnavailableHandler decorator.CommandHandler[MakeHoursUnavailable]

type makeHoursUnavailableHandler struct {
	hourRepo hour.Repository
}

func NewMakeHoursUnavailableHandler(
	hourRepo hour.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) MakeHoursUnavailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil") // TODO 开除预警
	}

	return decorator.ApplyCommandDecorators[MakeHoursUnavailable](
		makeHoursUnavailableHandler{hourRepo: hourRepo},
		logger,
		metricsClient,
	)
}

func (c makeHoursUnavailableHandler) Handle(ctx context.Context, cmd MakeHoursUnavailable) error {
	for _, hourToUpdate := range cmd.Hours {
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
