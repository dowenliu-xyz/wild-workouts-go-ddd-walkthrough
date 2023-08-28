package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	pkg_errors "github.com/pkg/errors"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/auth"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/errors"
)

type trainingRepository interface {
	FindTrainingsForUser(ctx context.Context, user auth.User) ([]Training, error)
	AllTrainings(ctx context.Context) ([]Training, error)
	CreateTraining(ctx context.Context, training Training, createFn func() error) error
	CancelTraining(ctx context.Context, trainingUUID string, deleteFn func(Training) error) error
	RescheduleTraining(ctx context.Context, trainingUUID string, newTime time.Time, updateFn func(Training) (Training, error)) error
	ApproveTrainingReschedule(ctx context.Context, trainingUUID string, updateFn func(Training) (Training, error)) error
	RejectTrainingReschedule(ctx context.Context, trainingUUID string, updateFn func(Training) (Training, error)) error
}

type userService interface {
	UpdateTrainingBalance(ctx context.Context, userID string, amountChange int) error
}

type trainerService interface {
	ScheduleTraining(ctx context.Context, trainingTime time.Time) error
	CancelTraining(ctx context.Context, trainingTime time.Time) error
}

type TrainingService struct {
	repo           trainingRepository
	trainerService trainerService
	userService    userService
}

func NewTrainingsService(
	repo trainingRepository,
	trainerService trainerService,
	userService userService,
) TrainingService {
	if repo == nil {
		panic("missing trainingRepository") // TODO 开除预警
	}
	if trainerService == nil {
		panic("missing trainerService") // TODO 开除预警
	}
	if userService == nil {
		panic("missing userService") // TODO 开除预警
	}

	return TrainingService{
		repo:           repo,
		trainerService: trainerService,
		userService:    userService,
	}
}

func (c TrainingService) GetAllTrainings(ctx context.Context) ([]Training, error) {
	return c.repo.AllTrainings(ctx)
}

func (c TrainingService) GetTrainingsForUser(ctx context.Context, user auth.User) ([]Training, error) {
	return c.repo.FindTrainingsForUser(ctx, user)
}

func (c TrainingService) CreateTraining(ctx context.Context, user auth.User, trainingTime time.Time, notes string) error {
	// sanity check
	if len(notes) > 1000 {
		return errors.NewIncorrectInputError("Note too big", "note-too-big")
	}

	training := Training{
		UUID:     uuid.New().String(),
		UserUUID: user.UUID,
		User:     user.DisplayName,
		Notes:    notes,
		Time:     trainingTime,
	}

	// TODO 这里会不会有性能问题？事务里调用闭包，闭包里远程调用，这不会导致长事务么？异步事务+最终一致性？

	return c.repo.CreateTraining(ctx, training, func() error {
		err := c.userService.UpdateTrainingBalance(ctx, user.UUID, -1)
		if err != nil {
			return pkg_errors.Wrap(err, "unable to change trainings balance")
		}

		err = c.trainerService.ScheduleTraining(ctx, training.Time)
		if err != nil {
			// TODO 如果是这里出错，前面 UpdateTrainingBalance 应该要取消吧。
			return pkg_errors.Wrap(err, "unable to schedule training")
		}

		return nil
	})
}

func (c TrainingService) RescheduleTraining(ctx context.Context, user auth.User, trainingUUID string, newTime time.Time, newNotes string) error {
	// sanity check
	if len(newNotes) > 1000 {
		return errors.NewIncorrectInputError("Note too big", "note-too-big")
	}

	// TODO 同样的事务中远程调用

	return c.repo.RescheduleTraining(ctx, trainingUUID, newTime, func(training Training) (Training, error) {
		if training.CanBeCancelled() {
			err := c.trainerService.ScheduleTraining(ctx, newTime)
			if err != nil {
				return Training{}, pkg_errors.Wrap(err, "unable to schedule training")
			}

			err = c.trainerService.CancelTraining(ctx, training.Time)
			if err != nil {
				// TODO 如果是这里出错，前面 ScheduleTraining 应该要取消吧。
				return Training{}, pkg_errors.Wrap(err, "unable to cancel training")
			}

			training.Time = newTime
			training.Notes = newNotes
		} else {
			training.ProposedTime = &newTime
			training.MoveProposedBy = &user.Role
			training.Notes = newNotes
		}

		return training, nil
	})
}

// TODO 好像 approve 之后，没有更新 hour 的 availability、trainer 的 schedule

func (c TrainingService) ApproveTrainingReschedule(ctx context.Context, user auth.User, trainingUUID string) error {
	return c.repo.ApproveTrainingReschedule(ctx, trainingUUID, func(training Training) (Training, error) {
		if training.ProposedTime == nil {
			return Training{}, pkg_errors.New("training has no proposed time")
		}
		if training.MoveProposedBy == nil {
			return Training{}, pkg_errors.New("training has no MoveProposedBy")
		}
		if *training.MoveProposedBy == "trainer" && training.UserUUID != user.UUID {
			return Training{}, pkg_errors.Errorf("user '%s' cannot approve reschedule of user '%s'", user.UUID, training.UserUUID)
		}
		if *training.MoveProposedBy == user.Role {
			return Training{}, pkg_errors.New("reschedule cannot be accepted by requesting person")
		}

		training.Time = *training.ProposedTime
		training.ProposedTime = nil

		return training, nil
	})
}

func (c TrainingService) RejectTrainingReschedule(ctx context.Context, user auth.User, trainingUUID string) error {
	return c.repo.RejectTrainingReschedule(ctx, trainingUUID, func(training Training) (Training, error) {
		if training.MoveProposedBy == nil {
			return Training{}, pkg_errors.New("training has no MoveProposedBy")
		}
		if *training.MoveProposedBy != "trainer" && training.UserUUID != user.UUID {
			return Training{}, pkg_errors.Errorf("user '%s' cannot approve reschedule of user '%s'", user.UUID, training.UserUUID)
		}

		training.ProposedTime = nil

		return training, nil
	})
}

func (c TrainingService) CancelTraining(ctx context.Context, user auth.User, trainingUUID string) error {
	return c.repo.CancelTraining(ctx, trainingUUID, func(training Training) error {
		if user.Role != "trainer" && training.UserUUID != user.UUID {
			return pkg_errors.Errorf("user '%s' is trying to cancel training of user '%s'", user.UUID, training.UserUUID)
		}

		var trainingBalanceDelta int
		if training.CanBeCancelled() {
			// just give training back
			trainingBalanceDelta = 1
		} else {
			if user.Role == "trainer" {
				// 1 for cancelled training +1 fine for cancelling by trainer less than 24h before training
				trainingBalanceDelta = 2
			} else {
				// fine for cancelling less than 24h before training
				trainingBalanceDelta = 0
			}
		}

		if trainingBalanceDelta != 0 {
			// TODO 事务中远程调用。。。
			err := c.userService.UpdateTrainingBalance(ctx, training.UserUUID, trainingBalanceDelta)
			if err != nil {
				return pkg_errors.Wrap(err, "unable to change trainings balance")
			}
		}

		err := c.trainerService.CancelTraining(ctx, training.Time)
		if err != nil {
			return pkg_errors.Wrap(err, "unable to cancel training")
		}

		return nil
	})
}
