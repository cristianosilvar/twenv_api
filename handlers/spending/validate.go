package spending

import (
	"fmt"
	"twenv/models"
)

func errParamIsRequired(name string) error {
	return fmt.Errorf("param: %s is required", name)
}

func validateSpending(spending *models.Spending) error {
	if spending.Value == 0 {
		return errParamIsRequired("value")
	}
	/* if spending.Date != time.Now() {
		return errParamIsRequired("date")
	} */
	return nil
}

func validateSpendingUpdate(spending *models.SpendingUpdate) error {
	if spending.Value == 0 && spending.Id == "" {
		return fmt.Errorf("request body is empty")
	}
	if spending.Id == "" {
		return errParamIsRequired("id")
	}
	if spending.Value == 0 {
		return errParamIsRequired("value")
	}
	/* if spending.Date != time.Now() {
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
