package schemas

import (
	"time"
)

type Spending struct {
	Value       float64
	Description string
	Date        time.Time
}
