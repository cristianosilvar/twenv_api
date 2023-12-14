package models

import (
	"time"
)

type Earning struct {
	Value       float64
	Description string
	Date        time.Time
}

type EarningResponse struct {
	Id          any       `bson:"_id,omitempty"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
