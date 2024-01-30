package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Spending struct {
	Value       float64
	Description string
	Date        time.Time
}

type SpendingResponse struct {
	Id          any       `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type SpendingUpdate struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Value       float64            `json:"value"`
	Description string             `json:"description"`
	Date        time.Time          `json:"date"`
}
