package repository

import "github.com/lustpills/wb_task0/pkg/repository"

type Orders interface {
}

type Repository struct {
	Orders
}

func NewService(repos *repository.Repository) *Repository {
	return &Repository{}
}
