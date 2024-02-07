package earning

import (
	"fmt"
	"twenv/models"
)

func errParamIsRequired(name string) error {
	return fmt.Errorf("param: %s is required", name)
}

func validateEarning(earning *models.Earning) error {
	if earning.Value == 0 {
		return errParamIsRequired("value")
	}
	/* if earning.Date != time.Now() {
		return errParamIsRequired("date")
	} */
	return nil
}

func validateEarningUpdate(earning *models.EarningUpdate) error {
	if earning.Value == 0 && earning.Id == "" {
		return fmt.Errorf("request body is empty")
	}
	if earning.Id == "" {
		return errParamIsRequired("id")
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
