package controller

import (
	"context"
	"niltasks/protoc"
)

type Service interface {
	GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*protoc.GetItemsResponse, error)
	CompleteItem(ctx context.Context, req *protoc.CompleteItemRequest) (*protoc.CompleteItemResponse, error)
	CreateItem(ctx context.Context, req *protoc.CreateItemRequest) (*protoc.CreateItemResponse, error)
	RescheduleItem(ctx context.Context, req *protoc.RescheduleItemRequest) (*protoc.RescheduleItemReponse, error)
	RemoveItem(ctx context.Context, req *protoc.RemoveItemRequest) (*protoc.RemoveItemResponse, error)
}

type Controller struct {
	service Service
}

func New(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*protoc.GetItemsResponse, error) {
	res, err := c.service.GetItems(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Controller) CompleteItem(ctx context.Context, req *protoc.CompleteItemRequest) (*protoc.CompleteItemResponse, error) {
	res, err := c.service.CompleteItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Controller) CreateItem(ctx context.Context, req *protoc.CreateItemRequest) (*protoc.CreateItemResponse, error) {
	res, err := c.service.CreateItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Controller) RescheduleItem(ctx context.Context, req *protoc.RescheduleItemRequest) (*protoc.RescheduleItemReponse, error) {
	res, err := c.service.RescheduleItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *Controller) RemoveItem(ctx context.Context, req *protoc.RemoveItemRequest) (*protoc.RemoveItemResponse, error) {
	res, err := c.service.RemoveItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
