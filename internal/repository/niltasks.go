package repository

import "fmt"

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) GetList() {
	fmt.Println("Hello from GO")
}
