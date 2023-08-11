package service

import (
	"context"
	"niltasks/internal/models"
	"niltasks/protoc"
)

type Repository interface {
	GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*models.List, error)
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*protoc.GetItemsResponse, error) {
	list, err := s.repo.GetItems(ctx, req)
	if err != nil {
		return nil, err
	}

	var res []*protoc.ToDoItem

	for _, item := range list.Tasks {
		res = append(res, &protoc.ToDoItem{
			Id:               item.Id,
			Title:            item.Title,
			Description:      item.Description,
			Completed:        item.Completed,
			ReadOnly:         item.ReadOnly,
			RescheduledTimes: item.RescheduledTimes,
			CreatedAt:        int32(item.CreatedAt.Unix()),
		})
	}

	return &protoc.GetItemsResponse{
		List: res,
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
