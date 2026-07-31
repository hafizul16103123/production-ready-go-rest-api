package logger

import (
	"log/slog"
	"os"
)

/*

 */

func Init() {

	// logger:=slog.New(
	// 	// slog.NewTextHandler(os.Stdout,nil), // Easy for humans to read.
	// 	slog.NewJSONHandler(os.Stdout,nil), // Perfect for:Elasticsearch Loki Datadog Cloud Logging
	// )

	// // Log to a specific file
	// logFile, _ := os.Create("api.log")
	// logger:= slog.New(slog.NewJSONHandler(logFile,&slog.HandlerOptions{AddSource:true}))

	log := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			},
		),
	)

	slog.SetDefault(log)
}
