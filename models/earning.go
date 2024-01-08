package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Earning struct {
	Value       float64
	Description string
	Date        time.Time
}

type EarningResponse struct {
	Id          interface{} `bson:"_id,omitempty"`
	Value       float64     `json:"value"`
	Description string      `json:"description"`
	Date        time.Time   `json:"date"`
}

type EarningUpdate struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Value       float64            `json:"value"`
	Description string             `json:"description"`
	Date        time.Time          `json:"date"`
}
