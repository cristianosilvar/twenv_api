package models

// import "time"

type User struct {
	Username string
	Email    string
	Password string
}

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
	// DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type CreateUserResponse struct {
	Id       interface{} `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
