package earning

import (
	"fmt"
	"twenv/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	/* if earning.Date != time.Now() {
		return errParamIsRequired("date")
	} */
	return nil
}

func validateEarningUpdate(earning *models.EarningUpdate) error {
	if earning.Description == "" && earning.Value == 0 && earning.Id == primitive.NilObjectID {
		return fmt.Errorf("request body is empty")
	}
	if earning.Id == primitive.NilObjectID {
		return errParamIsRequired("id")
	}
	if earning.Description == "" {
		return errParamIsRequired("description")
	}
	if earning.Value == 0 {
		return errParamIsRequired("value")
	}
	/* if earning.Date != time.Now() {
		return errParamIsRequired("date")
	} */
	return nil
}

func validateDelete(paramString string) error {
	if paramString == "" {
		return errParamIsRequired("id")
	}
	return nil
}
