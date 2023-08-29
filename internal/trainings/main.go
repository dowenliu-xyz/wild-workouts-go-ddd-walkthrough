package main

import (
	"context"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/chi/v5"

	grpcClient "github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/client"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/logs"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/server"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/adapters"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/app"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/app/command"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/app/query"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/ports"
)

func main() {
	logs.Init()

	ctx := context.Background()

	app, cleanup := newApplication(ctx)
	defer cleanup()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}

func newApplication(ctx context.Context) (app.Application, func()) {
	client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		panic(err)
	}

	trainerClient, closeTrainerClient, err := grpcClient.NewTrainerClient()
	if err != nil {
		panic(err)
	}

	usersClient, closeUsersClient, err := grpcClient.NewUsersClient()
	if err != nil {
		panic(err)
	}

	trainingsRepository := adapters.NewTrainingsFirestoreRepository(client)
	trainerGrpc := adapters.NewTrainerGrpc(trainerClient)
	usersGrpc := adapters.NewUsersGrpc(usersClient)

	return app.Application{
			Commands: app.Commands{
				ApproveTrainingReschedule: command.NewApproveTrainingRescheduleHandler(trainingsRepository, usersGrpc, trainerGrpc),
				CancelTraining:            command.NewCancelTrainingHandler(trainingsRepository, usersGrpc, trainerGrpc),
				RejectTrainingReschedule:  command.NewRejectTrainingRescheduleHandler(trainingsRepository),
				RescheduleTraining:        command.NewRescheduleTrainingHandler(trainingsRepository, usersGrpc, trainerGrpc),
				RequestTrainingReschedule: command.NewRequestTrainingRescheduleHandler(trainingsRepository),
				ScheduleTraining:          command.NewScheduleTrainingHandler(trainingsRepository, usersGrpc, trainerGrpc),
			},
			Queries: app.Queries{
				AllTrainings:     query.NewAllTrainingsHandler(trainingsRepository),
				TrainingsForUser: query.NewTrainingsForUserHandler(trainingsRepository),
			},
		}, func() {
			_ = closeTrainerClient()
			_ = closeUsersClient()
		}
}
