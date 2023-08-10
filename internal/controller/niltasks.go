package controller

import (
	"context"
	"niltasks/protoc"
)

type Service interface {
	GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*protoc.GetItemsResponse, error)
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

func (c *Controller) CompleteItem(context.Context, *protoc.CompleteItemRequest) (*protoc.CompleteItemResponse, error) {
	return &protoc.CompleteItemResponse{
		Id:        "test-id",
		Completed: true,
	}, nil
}

func (c *Controller) CreateItem(context.Context, *protoc.CreateItemRequest) (*protoc.CreateItemResponse, error) {
	return &protoc.CreateItemResponse{
		Item: &protoc.ToDoItem{
			Id:               "test-id",
			Title:            "Title",
			Description:      "Description",
			Completed:        false,
			ReadOnly:         false,
			RescheduledTimes: 0,
			CreatedAt:        1691675220,
		},
	}, nil
}

func (c *Controller) RescheduleItem(context.Context, *protoc.RescheduleItemRequest) (*protoc.RescheduleItemReponse, error) {
	return &protoc.RescheduleItemReponse{
		Id:               "test-id",
		RescheduledTimes: 1,
	}, nil
}

func (c *Controller) RemoveItem(context.Context, *protoc.RemoveItemRequest) (*protoc.RemoveItemResponse, error) {
	return &protoc.RemoveItemResponse{
		Id:      "test-id",
		Removed: true,
	}, nil
}
