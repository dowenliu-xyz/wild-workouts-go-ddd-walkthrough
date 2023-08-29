package query

import (
	"context"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/auth"
)

type TrainingsForUserHandler struct {
	readModel TrainingsForUserReadModel
}

func NewTrainingsForUserHandler(readModel TrainingsForUserReadModel) TrainingsForUserHandler {
	if readModel == nil {
		panic("nil readModel") // TODO 开除预警
	}

	return TrainingsForUserHandler{readModel: readModel}
}

type TrainingsForUserReadModel interface {
	FindTrainingsForUser(ctx context.Context, userUUID string) ([]Training, error)
}

func (h TrainingsForUserHandler) Handle(ctx context.Context, user auth.User) (tr []Training, err error) {
	return h.readModel.FindTrainingsForUser(ctx, user.UUID)
}
