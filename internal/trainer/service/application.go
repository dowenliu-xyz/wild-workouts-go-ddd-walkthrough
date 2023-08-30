package service

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/adapters"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/app"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/app/command"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/app/query"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/domain/hour"
)

func NewApplication(ctx context.Context) app.Application {
	firestoreClient, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		panic(err) // TODO 开除预警
	}

	factoryConfig := hour.FactoryConfig{
		MaxWeeksInTheFutureToSet: 6,
		MinUtcHour:               12,
		MaxUtcHour:               20,
	}

	datesRepository := adapters.NewDatesFirestoreRepository(firestoreClient, factoryConfig)

	hourFactory, err := hour.NewFactory(factoryConfig)
	if err != nil {
		panic(err) // TODO 开除预警
	}

	hourRepository := adapters.NewFirestoreHourRepository(firestoreClient, hourFactory)

	return app.Application{
		Commands: app.Commands{
			CancelTraining:       command.NewCancelTrainingHandler(hourRepository),
			ScheduleTraining:     command.NewScheduleTrainingHandler(hourRepository),
			MakeHoursAvailable:   command.NewMakeHoursAvailableHandler(hourRepository),
			MakeHoursUnavailable: command.NewMakeHoursUnavailableHandler(hourRepository),
		},
		Queries: app.Queries{
			HourAvailability:      query.NewHourAvailabilityHandler(hourRepository),
			TrainerAvailableHours: query.NewAvailableHoursHandler(datesRepository),
		},
	}
}
