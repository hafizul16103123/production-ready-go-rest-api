package logger
/*
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")

	slog.LevelDebug :will log all levels (Debug, Info, Warn, Error)
	slog.LevelInfo  :will log Info, Warn, Error (Debug will be ignored)
	slog.LevelWarn  :will log Warn, Error (Debug, Info will be ignored)
	slog.LevelError :will log only Error (Debug, Info, Warn will be ignored)
*/

import (
	"log/slog"
	"os"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/config"
)

func Init() {
	config:= config.Get()

	level := slog.LevelDebug

	if config.Env == "prod" {
		level = slog.LevelWarn
	}

	slog.Info("Info", "env", config.Env, "level", level)
	log := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     level,
			},
		),
	)

	slog.SetDefault(log)
}
