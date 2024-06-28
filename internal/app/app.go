package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/paniccaaa/sso/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, dbURI string, tokenTTL time.Duration) *App {
	// init db

	// init auth service

	// init grpc

	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
