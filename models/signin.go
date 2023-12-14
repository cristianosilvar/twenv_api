package models

type CreateSignRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
