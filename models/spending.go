package models

import (
	"time"
)

type Spending struct {
	Id          string
	UserID      interface{}
	Value       float64
	Description string
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type SpendingResponse struct {
	Id          any       `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}

type SpendingUpdate struct {
	Id          string    `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
