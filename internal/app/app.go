package app

import (
	"log/slog"
	"time"

	grpcapp "grpc-service-ref/internal/app/grpc"
	"grpc-service-ref/internal/services/auth"
	"grpc-service-ref/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	gRPCPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: initialize storage
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	// TODO: initialize auth service
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, gRPCPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
