package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/paniccaaa/sso/internal/app/grpc"
	"github.com/paniccaaa/sso/internal/services/auth"
	"github.com/paniccaaa/sso/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, dbURI string, tokenTTL time.Duration) *App {
	// init db
	storage, err := postgres.NewStorage(dbURI)
	if err != nil {
		panic(err)
	}

	// init auth service
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	// init grpc
	grpcApp := grpcapp.NewApp(log, grpcPort, authService, storage)

	return &App{
		GRPCServer: grpcApp,
	}
}
