package query

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/auth"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/decorator"
)

type TrainingsForUser struct {
	User auth.User
}

type TrainingsForUserHandler decorator.QueryHandler[TrainingsForUser, []Training]

type trainingsForUserHandler struct {
	readModel TrainingsForUserReadModel
}

func NewTrainingsForUserHandler(
	readModel TrainingsForUserReadModel,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) TrainingsForUserHandler {
	if readModel == nil {
		panic("nil readModel") // TODO 开除预警
	}

	return decorator.ApplyQueryDecorators[TrainingsForUser, []Training](
		trainingsForUserHandler{readModel: readModel},
		logger,
		metricsClient,
	)
}

type TrainingsForUserReadModel interface {
	FindTrainingsForUser(ctx context.Context, userUUID string) ([]Training, error)
}

func (h trainingsForUserHandler) Handle(ctx context.Context, query TrainingsForUser) (tr []Training, err error) {
	return h.readModel.FindTrainingsForUser(ctx, query.User.UUID)
}
