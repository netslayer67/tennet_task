package handlers

import (
	"task/repositories"
)

type handler struct {
	repo repositories.Repository
}

func NewHandler(repo repositories.Repository) *handler {
	return &handler{
		repo: repo,
	}
}
