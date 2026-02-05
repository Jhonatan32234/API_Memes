package users

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string
	CreatedAt time.Time
}

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
