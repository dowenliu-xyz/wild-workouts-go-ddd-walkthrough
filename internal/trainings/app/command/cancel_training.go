package command

import (
	"context"

	"github.com/pkg/errors"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/logs"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/domain/training"
)

type CancelTraining struct {
	TrainingUUID string
	User         training.User
}

type CancelTrainingHandler struct {
	repo           training.Repository
	userService    UserService
	trainerService TrainerService
}

func NewCancelTrainingHandler(repo training.Repository, userService UserService, trainerService TrainerService) CancelTrainingHandler {
	if repo == nil {
		panic("nil repo") // TODO 开除预警
	}
	if userService == nil {
		panic("nil user service") // 开除预警
	}
	if trainerService == nil {
		panic("nil trainer service")
	}

	return CancelTrainingHandler{repo: repo, userService: userService, trainerService: trainerService}
}

func (h CancelTrainingHandler) Handle(ctx context.Context, cmd CancelTraining) (err error) {
	defer func() {
		logs.LogCommandExecution("CancelTrainingHandler", cmd, err)
	}()

	return h.repo.UpdateTraining(
		ctx,
		cmd.TrainingUUID,
		cmd.User,
		func(ctx context.Context, tr *training.Training) (*training.Training, error) {
			if err := tr.Cancel(); err != nil {
				return nil, err
			}

			if balanceDelta := training.CancelBalanceDelta(*tr, cmd.User.Type()); balanceDelta != 0 {
				// TODO 事务中远程调用
				err := h.userService.UpdateTrainingBalance(ctx, tr.UserUUID(), balanceDelta)
				if err != nil {
					return nil, errors.WithMessage(err, "unable to change trainings balance")
				}
			}

			if err := h.trainerService.CancelTraining(ctx, tr.Time()); err != nil {
				return nil, errors.WithMessage(err, "unable to cancel training")
			}

			return tr, nil
		},
	)
}
