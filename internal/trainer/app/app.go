package app

import (
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/app/command"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CancelTraining   command.CancelTrainingHandler
	ScheduleTraining command.ScheduleTrainingHandler

	MakeHoursAvailable   command.MakeHoursAvailableHandler
	MakeHoursUnavailable command.MakeHoursUnavailableHandler
}

type Queries struct {
	HourAvailability      query.HourAvailabilityHandler
	TrainerAvailableHours query.AvailableHoursHandler
}
