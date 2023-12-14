package models

import (
	"time"
)

type Spending struct {
	Value       float64
	Description string
	Date        time.Time
}

type SpendingResponse struct {
	Id          any       `bson:"_id,omitempty"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
