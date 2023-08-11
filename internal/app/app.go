package app

import (
	"fmt"
	"net/http"
	"niltasks/config"
	"niltasks/internal/controller"
	"niltasks/internal/repository"
	"niltasks/internal/service"
	"niltasks/pkg/mongo"
	"niltasks/protoc"
)

func Serve() {
	cfg := config.MustLoad()

	mongo := mongo.New(&cfg.Mongo)
	fmt.Println("MongoDB connected")

	repo := repository.New(cfg, mongo)
	service := service.New(repo)
	controller := controller.New(service)

	twirpHandler := protoc.NewToDoItemsServer(controller)

	host := cfg.Server.Host + ":" + cfg.Server.Port
	fmt.Printf("Starting server on %s\n", host)
	err := http.ListenAndServe(host, twirpHandler)
	if err != nil {
		panic(err)
	}
}
