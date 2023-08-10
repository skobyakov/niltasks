package repository

import (
	"fmt"
	"niltasks/config"
	"niltasks/pkg/mongo"
)

type Repository struct {
	db *mongo.MongoDB
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		db: mongo.New(&cfg.Mongo),
	}
}

func (r *Repository) GetList() {
	fmt.Println("Hello from GO")
}
