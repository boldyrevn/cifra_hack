package model

type User struct {
	Name    string `json:"name"`
	Role    string `json:"role"`
	Balance int    `json:"balance"`
	Email   string `json:"email"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
