package app

import (
	"niltasks/internal/controller"
	"niltasks/internal/repository"
	"niltasks/internal/service"
)

func Serve() {
	repo := repository.New()
	service := service.New(repo)
	controller := controller.New(service)

	controller.GetList()
}
