package main

import (
	"context"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/chi/v5"

	client_pkg "github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/client"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/logs"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/server"
)

func main() {
	logs.Init()

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		panic(err)
	}

	trainerClient, closeTrainerClient, err := client_pkg.NewTrainerClient()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = closeTrainerClient()
	}()

	usersClient, closeUsersClient, err := client_pkg.NewUsersClient()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = closeUsersClient()
	}()

	firebaseDB := db{client}

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return HandlerFromMux(HttpServer{firebaseDB, trainerClient, usersClient}, router)
	})
}
