package command

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/decorator"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/errors"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/domain/hour"
)

type CancelTraining struct {
	Hour time.Time
}

type CancelTrainingHandler decorator.CommandHandler[CancelTraining]

type cancelTrainingHandler struct {
	hourRepo hour.Repository
}

func NewCancelTrainingHandler(
	hourRepo hour.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CancelTrainingHandler {
	if hourRepo == nil {
		panic("nil hourRepo") // TODO 开除预警
	}

	return decorator.ApplyCommandDecorators[CancelTraining](
		cancelTrainingHandler{hourRepo: hourRepo},
		logger,
		metricsClient,
	)
}

func (h cancelTrainingHandler) Handle(ctx context.Context, cmd CancelTraining) error {
	if err := h.hourRepo.UpdateHour(ctx, cmd.Hour, func(h *hour.Hour) (*hour.Hour, error) {
		if err := h.CancelTraining(); err != nil {
			return nil, err
		}
		return h, nil
	}); err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-availability")
	}

	return nil
}
