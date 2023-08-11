package repository

import (
	"context"
	"niltasks/config"
	"niltasks/internal/models"
	"niltasks/pkg/mongo"
	"niltasks/protoc"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	listsColection = "lists"
)

type Repository struct {
	db *mongo.MongoDB
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		db: mongo.New(&cfg.Mongo),
	}
}

func (r *Repository) GetItems(ctx context.Context, req *protoc.GetItemsRequest) (*models.List, error) {
	coll := r.db.Database.Collection(listsColection)

	var res *models.List

	err := coll.FindOne(ctx, bson.D{{Key: "user_id", Value: req.GetUserId()}}).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
