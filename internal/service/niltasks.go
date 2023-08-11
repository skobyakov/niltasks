package service

import (
	"context"
	"niltasks/protoc"
)

type Repository interface {
	GetList()
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*protoc.GetItemsResponse, error) {
	list := []*protoc.ToDoItem{
		{
			Id:               "test-id",
			Title:            "Title",
			Description:      "Description",
			Completed:        false,
			ReadOnly:         false,
			RescheduledTimes: 1,
			CreatedAt:        1691675220,
		},
	}
	return &protoc.GetItemsResponse{
		List: list,
	}, nil
}

func (s *Service) CompleteItem(ctx context.Context, req *protoc.CompleteItemRequest) (*protoc.CompleteItemResponse, error) {
	return &protoc.CompleteItemResponse{
		Id:        "test-id",
		Completed: true,
	}, nil
}

func (s *Service) CreateItem(ctx context.Context, req *protoc.CreateItemRequest) (*protoc.CreateItemResponse, error) {
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

func (s *Service) RescheduleItem(ctx context.Context, req *protoc.RescheduleItemRequest) (*protoc.RescheduleItemReponse, error) {
	return &protoc.RescheduleItemReponse{
		Id:               "test-id",
		RescheduledTimes: 1,
	}, nil
}

func (s *Service) RemoveItem(ctx context.Context, req *protoc.RemoveItemRequest) (*protoc.RemoveItemResponse, error) {
	return &protoc.RemoveItemResponse{
		Id:      "test-id",
		Removed: true,
	}, nil
}
