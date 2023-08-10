package mongo

import (
	"context"
	"fmt"
	"niltasks/pkg/mongo/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Conn     *mongo.Client
	Database *mongo.Database
}

func mustConnect(cfg *config.MongoConfig) *mongo.Client {
	url := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	var ping bson.M
	if err := client.Database(cfg.Database).RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&ping); err != nil {
		panic(err)
	}

	return client
}

func New(cfg *config.MongoConfig) *MongoDB {
	client := mustConnect(cfg)

	return &MongoDB{
		Conn:     client,
		Database: client.Database(cfg.Database),
	}
}
