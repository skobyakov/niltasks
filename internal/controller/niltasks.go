package controller

import (
	"context"
	"fmt"
	"niltasks/protoc"
)

type Service interface {
	GetList()
}

type Controller struct {
	service Service
}

func New(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) GetList() {
	c.service.GetList()
}

func (c *Controller) GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*protoc.GetItemsResponse, error) {
	fmt.Println(req.GetUserId())
	c.service.GetList()
	list := []*protoc.ToDoItem{
		&protoc.ToDoItem{Title: "Hello Twirp", Active: false, Completed: false},
	}
	return &protoc.GetItemsResponse{
		List: list,
	}, nil
}
