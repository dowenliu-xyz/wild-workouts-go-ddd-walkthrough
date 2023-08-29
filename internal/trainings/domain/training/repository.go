package training

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type NotFoundError struct {
	TrainingUUID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("training '%s' not found", e.TrainingUUID)
}

func NewNotFoundError(trainingUUID string) error {
	return errors.WithStack(NotFoundError{TrainingUUID: trainingUUID})
}

type Repository interface {
	AddTraining(ctx context.Context, tr *Training) error

	GetTraining(ctx context.Context, trainingUUID string, user User) (*Training, error)

	UpdateTraining(
		ctx context.Context,
		trainingUUID string,
		user User,
		updateFn func(ctx context.Context, tr *Training) (*Training, error),
	) error
}
