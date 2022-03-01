package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	CreatorId   string             `json:"creatorId" validate:"required"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description,omitempty"`
	Date        time.Time          `json:"date"`
	Completed   bool               `json:"completed"`
}
