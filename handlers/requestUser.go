package handlers

import (
	"fmt"
	"time"
)

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

type Spending struct {
	Value       float64
	Description string
	Date        time.Time
}

type Delete struct {
	Id string `json:"id"`
}

type SpendingResponse struct {
	Id          any       `bson:"_id,omitempty"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (spending *Spending) ValidateSpending() error {
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

func (item *Delete) ValidateDelete() error {
	if item.Id == "" {
		return errParamIsRequired("id")
	}
	return nil
}
