package models

import "time"

type Task struct {
	Id               string    `bson:"_id"`
	Title            string    `bson:"title"`
	Description      string    `bson:"description"`
	CreatedAt        time.Time `bson:"created_at"`
	Completed        bool      `bson:"completed"`
	ReadOnly         bool      `bson:"readOnly"`
	RescheduledTimes int32     `bson:"rescheduledTimes"`
}

type List struct {
	Id     string `bson:"_id"`
	UserId string `bson:"user_id"`
	Tasks  []Task `bson:"tasks"`
}
