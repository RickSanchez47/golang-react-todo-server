package models

import (
	"go.mongodb.org/mongo-driver/mongo/bson/primitive"
)

type TodoList struct {
	ID     primitive.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"_task,omitempty"`
	Status bool               `json:"status"`
}
