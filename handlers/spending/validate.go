package spending

import (
	"fmt"
	"twenv/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func errParamIsRequired(name string) error {
	return fmt.Errorf("param: %s is required", name)
}

func validateSpending(spending *models.Spending) error {
	if spending.Description == "" && spending.Value == 0 {
		return fmt.Errorf("request body is empty")
	}
	if spending.Description == "" {
		return errParamIsRequired("description")
	}
	if spending.Value == 0 {
		return errParamIsRequired("value")
	}
	/* if spending.Date != time.Now() {
		return errParamIsRequired("date")
	} */
	return nil
}

func validateSpendingUpdate(spending *models.SpendingUpdate) error {
	if spending.Description == "" && spending.Value == 0 && spending.Id == primitive.NilObjectID {
		return fmt.Errorf("request body is empty")
	}
	if spending.Id == primitive.NilObjectID {
		return errParamIsRequired("id")
	}
	if spending.Description == "" {
		return errParamIsRequired("description")
	}
	if spending.Value == 0 {
		return errParamIsRequired("value")
	}
	/* if spending.Date != time.Now() {
		return errParamIsRequired("date")
	} */
	return nil
}

func validateDelete(item *models.Delete) error {
	if item.Id == primitive.NilObjectID {
		return errParamIsRequired("id")
	}
	return nil
}
