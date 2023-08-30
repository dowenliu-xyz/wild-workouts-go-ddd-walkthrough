package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/logs"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/server"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/ports"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/service"
)

func main() {
	logs.Init()

	ctx := context.Background()

	application, cleanup := service.NewApplication(ctx)
	defer cleanup()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(application), router)
	})
}
