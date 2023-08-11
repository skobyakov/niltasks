package service

import (
	"context"
	"niltasks/internal/models"
	"niltasks/protoc"
)

type Repository interface {
	GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*models.List, error)
	CreateItem(ctx context.Context, req *protoc.CreateItemRequest) (*models.Task, error)
	RescheduleItem(ctx context.Context, req *protoc.RescheduleItemRequest) (*models.Task, error)
	RemoveItem(ctx context.Context, req *protoc.RemoveItemRequest) error
	CompleteItem(ctx context.Context, req *protoc.CompleteItemRequest) (*models.Task, error)
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
			Id:               item.ID.Hex(),
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
	res, err := s.repo.CreateItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return &protoc.CreateItemResponse{
		Item: &protoc.ToDoItem{
			Id:               res.ID.Hex(),
			Title:            res.Title,
			Description:      res.Description,
			Completed:        false,
			ReadOnly:         false,
			RescheduledTimes: 0,
			CreatedAt:        int32(res.CreatedAt.Unix()),
		},
	}, nil
}

func (s *Service) RescheduleItem(ctx context.Context, req *protoc.RescheduleItemRequest) (*protoc.RescheduleItemReponse, error) {
	task, err := s.repo.RescheduleItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &protoc.RescheduleItemReponse{
		Id:               task.ID.Hex(),
		RescheduledTimes: task.RescheduledTimes,
	}, nil
}

func (s *Service) RemoveItem(ctx context.Context, req *protoc.RemoveItemRequest) (*protoc.RemoveItemResponse, error) {
	err := s.repo.RemoveItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &protoc.RemoveItemResponse{
		Id:      req.GetItemId(),
		Removed: true,
	}, nil
}
