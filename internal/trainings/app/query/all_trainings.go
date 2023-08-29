package query

import (
	"context"
)

type AllTrainingsHandler struct {
	readModel AllTrainingsReadModel
}

func NewAllTrainingsHandler(readModel AllTrainingsReadModel) AllTrainingsHandler {
	if readModel == nil {
		panic("nil readModel") // TODO 开除预警
	}

	return AllTrainingsHandler{readModel: readModel}
}

type AllTrainingsReadModel interface {
	AllTrainings(ctx context.Context) ([]Training, error)
}

func (h AllTrainingsHandler) Handle(ctx context.Context) (tr []Training, err error) {
	return h.readModel.AllTrainings(ctx)
}
