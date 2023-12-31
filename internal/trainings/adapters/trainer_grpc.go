package adapters

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/genproto/trainer"
)

type TrainerGrpc struct {
	client trainer.TrainerServiceClient
}

func NewTrainerGrpc(client trainer.TrainerServiceClient) TrainerGrpc {
	return TrainerGrpc{client: client}
}

func (s TrainerGrpc) ScheduleTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.client.ScheduleTraining(ctx, &trainer.UpdateHourRequest{
		Time: timestamppb.New(trainingTime),
	})

	return errors.WithStack(err)
}

func (s TrainerGrpc) CancelTraining(ctx context.Context, trainingTime time.Time) error {
	_, err := s.client.CancelTraining(ctx, &trainer.UpdateHourRequest{
		Time: timestamppb.New(trainingTime),
	})

	return errors.WithStack(err)
}

func (s TrainerGrpc) MoveTraining(
	ctx context.Context,
	newTime time.Time,
	originalTrainingTime time.Time,
) error {
	err := s.ScheduleTraining(ctx, newTime)
	if err != nil {
		return errors.WithMessage(err, "unable to schedule training")
	}

	err = s.CancelTraining(ctx, originalTrainingTime)
	if err != nil {
		return errors.WithMessage(err, "unable to cancel training")
	}

	return nil
}
