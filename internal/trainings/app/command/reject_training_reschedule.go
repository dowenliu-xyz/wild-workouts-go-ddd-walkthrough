package command

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/decorator"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/logs"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/domain/training"
)

type RejectTrainingReschedule struct {
	TrainingUUID string
	User         training.User
}

type RejectTrainingRescheduleHandler decorator.CommandHandler[RejectTrainingReschedule]

type rejectTrainingRescheduleHandler struct {
	repo training.Repository
}

func NewRejectTrainingRescheduleHandler(
	repo training.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) RejectTrainingRescheduleHandler {
	if repo == nil {
		panic("nil repo service") // TODO 开除预警
	}

	return decorator.ApplyCommandDecorators[RejectTrainingReschedule](
		rejectTrainingRescheduleHandler{repo: repo},
		logger,
		metricsClient,
	)
}

func (h rejectTrainingRescheduleHandler) Handle(ctx context.Context, cmd RejectTrainingReschedule) (err error) {
	defer func() {
		logs.LogCommandExecution("RejectTrainingReschedule", cmd, err)
	}()

	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			if err := tr.RejectReschedule(); err != nil {
				return nil, err
			}

			return tr, nil
		},
	)
}
