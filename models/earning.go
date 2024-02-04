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
	Id          any       `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type EarningUpdate struct {
	Id          string    `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
