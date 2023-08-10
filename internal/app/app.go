package app

import (
	"fmt"
	"net/http"
	"niltasks/config"
	"niltasks/internal/controller"
	"niltasks/internal/repository"
	"niltasks/internal/service"
	"niltasks/protoc"
)

func Serve() {
	cfg := config.MustLoad()

	repo := repository.New(cfg)
	service := service.New(repo)
	controller := controller.New(service)

	twirpHandler := protoc.NewToDoItemsServer(controller)

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8080", twirpHandler)
	if err != nil {
		panic(err)
	}
}
