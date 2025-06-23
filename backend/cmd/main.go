package main

import (
	"e-shop-management-system/internal/config"
	"e-shop-management-system/internal/server"
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
