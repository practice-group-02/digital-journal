package server

import (
	"database/sql"
	"digital-journal/internal/config"
	"digital-journal/internal/dal"
	"digital-journal/internal/handlers"
	"fmt"
	"log/slog"
	"net/http"
)

type server struct {
	config *config.Config
	mux    *http.ServeMux
	store  *sql.DB
}

func NewServer(config *config.Config) *server {
	err := dal.NewStore(*config.DBconfigs)
	if err != nil {
		slog.Error("cannot connecto database!", "error", err)
	} else {
		slog.Info("successfully connected to database", "db port", config.DBconfigs.DBport)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", handlers.CreateUser)
	mux.HandleFunc("POST /login", handlers.LoginUser)
	mux.HandleFunc("POST /program", handlers.PostProgram)
	mux.HandleFunc("GET /programs", handlers.GetPrograms)
	mux.HandleFunc("GET /programs/tags", handlers.GetProgramsWithTags)
	mux.HandleFunc("GET /programs/{Type}", handlers.GetProgramsOfType)
	mux.HandleFunc("PUT /program/{Id}", handlers.UpdateProgram)
	mux.HandleFunc("DELETE /program/{Id}", handlers.DeleteProgram)

	return &server{
		mux:    mux,
		config: config,
		store:  dal.DB,
	}
}

func (s *server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", s.config.AppConfig.Port), s.mux)
}
