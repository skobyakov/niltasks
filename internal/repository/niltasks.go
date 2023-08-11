package repository

import (
	"context"
	"niltasks/config"
	"niltasks/internal/models"
	"niltasks/pkg/mongo"
	"niltasks/protoc"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	listsColection = "lists"
)

type Repository struct {
	db *mongo.MongoDB
}

func New(cfg *config.Config, db *mongo.MongoDB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*models.List, error) {
	coll := r.db.Database.Collection(listsColection)

	var res *models.List

	newList := models.List{
		UserId: req.UserId,
		Tasks:  []models.Task{},
	}

	find := bson.D{{Key: "user_id", Value: req.GetUserId()}}
	update := bson.D{{Key: "$setOnInsert", Value: newList}}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	err := coll.FindOneAndUpdate(ctx, find, update, opts).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) CreateItem(ctx context.Context, req *protoc.CreateItemRequest) (*protoc.CreateItemResponse, error) {
	coll := r.db.Database.Collection(listsColection)

	task := models.Task{
		Title:            req.GetTitle(),
		Description:      req.GetDescription(),
		CreatedAt:        time.Now(),
		Completed:        false,
		ReadOnly:         false,
		RescheduledTimes: 0,
	}
	find := bson.D{{Key: "user_id", Value: req.GetUserId()}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "tasks", Value: task}}}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := coll.FindOneAndUpdate(ctx, find, update, opts).Decode(&task)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
