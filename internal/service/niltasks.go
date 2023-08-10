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
