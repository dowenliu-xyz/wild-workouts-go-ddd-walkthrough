package command

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/decorator"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/logs"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/domain/training"
)

type RescheduleTraining struct {
	TrainingUUID string
	NewTime      time.Time

	User training.User

	NewNotes string
}

type RescheduleTrainingHandler decorator.CommandHandler[RescheduleTraining]

type rescheduleTrainingHandler struct {
	repo           training.Repository
	userService    UserService
	trainerService TrainerService
}

func NewRescheduleTrainingHandler(
	repo training.Repository,
	userService UserService,
	trainerService TrainerService,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) RescheduleTrainingHandler {
	if repo == nil {
		panic("nil repo") // TODO 开除预警
	}
	if userService == nil {
		panic("nil userService") // TODO 开除预警
	}
	if trainerService == nil {
		panic("nil trainerService") // TODO 开除预警
	}

	return decorator.ApplyCommandDecorators[RescheduleTraining](
		rescheduleTrainingHandler{repo: repo, userService: userService, trainerService: trainerService},
		logger,
		metricsClient,
	)
}

func (h rescheduleTrainingHandler) Handle(ctx context.Context, cmd RescheduleTraining) (err error) {
	defer func() {
		logs.LogCommandExecution("RescheduleTraining", cmd, err)
	}()

	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			originalTrainingTime := tr.Time()

			if err := tr.UpdateNotes(cmd.NewNotes); err != nil {
				return nil, err
			}

			if err := tr.RescheduleTraining(cmd.NewTime); err != nil {
				return nil, err
			}

			err := h.trainerService.MoveTraining(ctx, cmd.NewTime, originalTrainingTime)
			if err != nil {
				return nil, err
			}

			return tr, nil
		},
	)
}
