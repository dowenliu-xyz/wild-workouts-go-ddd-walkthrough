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
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/ports"
)

func main() {
	logs.Init()

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		panic(err)
	}

	trainerClient, closeTrainerClient, err := grpcClient.NewTrainerClient()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = closeTrainerClient()
	}()

	usersClient, closeUsersClient, err := grpcClient.NewUsersClient()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = closeUsersClient()
	}()

	trainingsRepository := adapters.NewTrainingsFirestoreRepository(client)
	trainerGrpc := adapters.NewTrainerGrpc(trainerClient)
	usersGrpc := adapters.NewUsersGrpc(usersClient)

	trainingsService := app.NewTrainingsService(trainingsRepository, trainerGrpc, usersGrpc)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(trainingsService), router)
	})
}
