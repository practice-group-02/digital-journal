package main

import (
	"digital-journal/internal/config"
	"digital-journal/internal/server"
	"log/slog"
)

func main() {
	config := config.NewConfig()
	s := server.NewServer(config)

	slog.Info("Server is running", "port", config.AppConfig.Port)

	if err := s.Run(); err != nil {
		slog.Error("error to run server", "error", err.Error())
	}
}
