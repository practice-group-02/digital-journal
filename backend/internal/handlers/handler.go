package handlers

import (
	"e-shop-management-system/internal/config"
	"e-shop-management-system/internal/dal"
)

type Handler struct {
	config           *config.Config
	store            *dal.Store
}

func NewHandler(config *config.Config, store *dal.Store) *Handler {
	return &Handler{
		config:	config,
		store: store,
	}
}