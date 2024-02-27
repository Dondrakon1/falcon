package main

import (
	"context"
	"falcon/internal/app"
	"falcon/internal/config"
	"falcon/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	ctx := context.WithoutCancel(context.Background())

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	go func() {
		<-shutdownChan
		log.Info("app.Shutdown")
		os.Exit(0)
	}()

	scorpion := app.NewApp(cfg, log)
	log.Info("app.Run", slog.String("env", cfg.Env))
	if err := scorpion.Run(); err != nil {
		log.Error("app.Run: %w", err)
		os.Exit(1)
	}
	<-ctx.Done()
	scorpion.Stop()

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
