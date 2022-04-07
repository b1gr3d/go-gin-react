package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Change struct {
	//TODO make specific items required
	ID   primitive.ObjectID `bson:"id"`
	User *string            `json:"user"`
	Env  *string            `json:"env"`
	App  *string            `json:"app"`
	Desc *string            `json:"desc"`
	Date *string            `json:"date"`
}
