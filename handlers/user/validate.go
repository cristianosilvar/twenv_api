package user

import (
	"fmt"
	"twenv/models"
)

func errParamIsRequired(name string) error {
	return fmt.Errorf("param: %s is required", name)
}

func validate(u *models.User) error {
	if u.Username == "" && u.Password == "" && u.Email == "" {
		return fmt.Errorf("request body is empty")
	}
	if u.Username == "" {
		return errParamIsRequired("username")
	}
	if u.Email == "" {
		return errParamIsRequired("email")
	}
	if u.Password == "" {
		return errParamIsRequired("password")
	}
	return nil
}

func ValidateSignIn(u *models.CreateSignRequest) error {
	if u.Password == "" && u.Email == "" {
		return fmt.Errorf("request body is empty")
	}
	if u.Email == "" {
		return errParamIsRequired("email")
	}
	if u.Password == "" {
		return errParamIsRequired("password")
	}
	return nil
}
