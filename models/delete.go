package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Delete struct {
	Id primitive.ObjectID `json:"id"`
}
