package spending

import (
	"fmt"
	"twenv/models"
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
	return nil
}

func validateDelete(item *models.Delete) error {
	if item.Id == "" {
		return errParamIsRequired("id")
	}
	return nil
}
