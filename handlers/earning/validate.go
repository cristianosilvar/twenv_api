package earning

import (
	"fmt"
	"twenv/models"
)

func errParamIsRequired(name string) error {
	return fmt.Errorf("param: %s is required", name)
}

func validateEarning(earning *models.Earning) error {
	if earning.Description == "" && earning.Value == 0 {
		return fmt.Errorf("request body is empty")
	}
	if earning.Description == "" {
		return errParamIsRequired("description")
	}
	if earning.Value == 0 {
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
