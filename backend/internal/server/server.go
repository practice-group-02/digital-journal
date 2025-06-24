package server

import (
	"digital-journal/internal/config"
	"digital-journal/internal/dal"
	"fmt"
	"log/slog"
	"net/http"
)

type server struct {
	config  *config.Config
	mux     *http.ServeMux
}

func NewServer(config *config.Config) *server {
	err := dal.NewStore(*config.DBconfigs)
	if err != nil {
		slog.Error("cannot connecto database!", "error", err)
	} else {
		slog.Info("successfully connected to database", "db port", config.DBconfigs.DBport)
	}

	return &server{
		mux:     http.NewServeMux(),
		config:  config,
	}
}

func (s *server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", s.config.AppConfig.Port), s.mux)
}