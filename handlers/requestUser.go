package handlers

import "fmt"

func errParamIsRequired(name string) error {
	return fmt.Errorf("param: %s is required", name)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *CreateUserRequest) Validate() error {
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

type CreateSignRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *CreateSignRequest) ValidateSignIn() error {
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
