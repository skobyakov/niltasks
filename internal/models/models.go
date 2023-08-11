package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Title            string             `bson:"title"`
	Description      string             `bson:"description"`
	CreatedAt        time.Time          `bson:"created_at"`
	Completed        bool               `bson:"completed"`
	ReadOnly         bool               `bson:"readOnly"`
	RescheduledTimes int32              `bson:"rescheduledTimes"`
}

type List struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId string             `bson:"user_id"`
	Tasks  []Task             `bson:"tasks"`
}
