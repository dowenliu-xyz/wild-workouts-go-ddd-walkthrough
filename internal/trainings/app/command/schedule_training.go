package command

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/logs"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/domain/training"
)

type ScheduleTraining struct {
	TrainingUUID string

	UserUUID string
	UserName string

	TrainingTime time.Time
	Notes        string
}

type ScheduleTrainingHandler struct {
	repo           training.Repository
	userService    UserService
	trainerService TrainerService
}

func NewScheduleTrainingHandler(repo training.Repository, userService UserService, trainerService TrainerService) ScheduleTrainingHandler {
	if repo == nil {
		panic("nil repo") // TODO 开除预警
	}
	if userService == nil {
		panic("nil userService") // TODO 开除预警
	}
	if trainerService == nil {
		panic("nil trainerService") // TODO 开除预警
	}

	return ScheduleTrainingHandler{repo: repo, userService: userService, trainerService: trainerService}
}

func (h ScheduleTrainingHandler) Handle(ctx context.Context, cmd ScheduleTraining) (err error) {
	defer func() {
		logs.LogCommandExecution("ScheduleTraining", cmd, err)
	}()

	tr, err := training.NewTraining(cmd.TrainingUUID, cmd.UserUUID, cmd.UserName, cmd.TrainingTime)
	if err != nil {
		return err
	}

	if err := h.repo.AddTraining(ctx, tr); err != nil {
		return err
	}

	err = h.userService.UpdateTrainingBalance(ctx, tr.UserUUID(), -1)
	if err != nil {
		return errors.Wrap(err, "unable to change trainings balance")
	}

	err = h.trainerService.ScheduleTraining(ctx, tr.Time())
	if err != nil {
		return errors.WithMessage(err, "unable to schedule training")
	}

	return nil
}
