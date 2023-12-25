package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/HAtherlolz/grpc/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.Server
}

func New(log *slog.Logger, gRPCPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: initialize storage

	// TODO: initialize auth service

	grpcApp := grpcapp.New(log, gRPCPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
