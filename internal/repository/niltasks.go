package repository

import (
	"context"
	"niltasks/config"
	"niltasks/internal/models"
	"niltasks/pkg/mongo"
	"niltasks/protoc"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *Repository) CreateItem(ctx context.Context, req *protoc.CreateItemRequest) (*models.Task, error) {
	coll := r.db.Database.Collection(listsColection)

	task := &models.Task{
		ID:               primitive.NewObjectIDFromTimestamp(time.Now()),
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

	err := coll.FindOneAndUpdate(ctx, find, update, opts).Decode(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Repository) RescheduleItem(ctx context.Context, req *protoc.RescheduleItemRequest) (*models.Task, error) {
	coll := r.db.Database.Collection(listsColection)

	var task models.Task

	taskId, err := primitive.ObjectIDFromHex(req.GetItemId())
	if err != nil {
		return nil, err
	}

	find := bson.D{{Key: "user_id", Value: req.GetUserId()}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "tasks.$[elem].rescheduledTimes", Value: 1}}}}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After).SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.D{{Key: "elem._id", Value: taskId}}},
	})

	err = coll.FindOneAndUpdate(ctx, find, update, opt).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *Repository) RemoveItem(ctx context.Context, req *protoc.RemoveItemRequest) error {
	coll := r.db.Database.Collection(listsColection)

	taskId, err := primitive.ObjectIDFromHex(req.GetItemId())
	if err != nil {
		return err
	}

	find := bson.D{{Key: "user_id", Value: req.GetUserId()}}
	update := bson.D{{Key: "$pull", Value: bson.D{{Key: "tasks", Value: bson.D{{Key: "_id", Value: taskId}}}}}}

	err = coll.FindOneAndUpdate(ctx, find, update).Err()

	return err
}

func (r *Repository) CompleteItem(ctx context.Context, req *protoc.CompleteItemRequest) (*models.Task, error) {
	coll := r.db.Database.Collection(listsColection)

	var task models.Task

	taskId, err := primitive.ObjectIDFromHex(req.GetItemId())
	if err != nil {
		return nil, err
	}
	// TODO: Bug - need to return task not list (or change contract)
	find := bson.D{{Key: "user_id", Value: req.GetUserId()}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "tasks.$[elem].completed", Value: true}}}}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After).SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.D{{Key: "elem._id", Value: taskId}}},
	})
	err = coll.FindOneAndUpdate(ctx, find, update, opt).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
