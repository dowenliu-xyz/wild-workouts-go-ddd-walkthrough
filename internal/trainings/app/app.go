package app

import (
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/app/command"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ApproveTrainingReschedule command.ApproveTrainingRescheduleHandler
	CancelTraining            command.CancelTrainingHandler
	RejectTrainingReschedule  command.RejectTrainingRescheduleHandler
	RescheduleTraining        command.RescheduleTrainingHandler
	RequestTrainingReschedule command.RequestTrainingRescheduleHandler
	ScheduleTraining          command.ScheduleTrainingHandler
}

type Queries struct {
	AllTrainings     query.AllTrainingsHandler
	TrainingsForUser query.TrainingsForUserHandler
}
