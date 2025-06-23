package server

import (
	"e-shop-management-system/internal/config"
	"e-shop-management-system/internal/dal"
	"e-shop-management-system/internal/handlers"
	"fmt"
	"log/slog"
	"net/http"
)

type server struct {
	config  *config.Config
	handler *handlers.Handler
	mux     *http.ServeMux
}

func NewServer(config *config.Config) *server {
	store, err := dal.NewStore(*config.DBconfigs)
	if err != nil {
		slog.Error("cannot connecto database!", "error", err)
	} else {
		slog.Info("successfully connected to database", "db port", config.DBconfigs.DBport)
	}

	return &server{
		mux:     http.NewServeMux(),
		config:  config,
		handler: handlers.NewHandler(config, store),
	}
}

func (s *server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", s.config.AppConfig.Port), s.mux)
}