package app

import (
	"falcon/internal/camera"
	"falcon/internal/config"
	"falcon/internal/service/code"
	"falcon/internal/storage/pg"
	"fmt"
	"log/slog"
)

type App struct {
	cfg     *config.Config
	log     *slog.Logger
	storage code.Storage
	service code.StorageService
	camera  codeListener
}

type codeListener interface {
	StartListening()
	Close()
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

func (a *App) Run() error {
	op := "app.Run"
	storage, err := pg.New(a.cfg.StoragePath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	service := code.NewCodeService(storage)
	service.Log = a.log
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	cam, err := camera.NewCamera(a.cfg.CameraAddress, service)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	cam.Log = a.log

	cam.StartListening()

	return nil
}

func (a *App) Stop() {
	//
	//a.camera.Disconnect()
	//a.storage.Disconnect()
	//a.service.Stop()

	a.log.Info("app.Stop", slog.String("env", a.cfg.Env))

}
