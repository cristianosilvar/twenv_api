package schemas

import "gorm.io/gorm"

type Spending struct {
	gorm.Model

	Value       float64
	Description string
	Date        string
}
