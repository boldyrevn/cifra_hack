package model

type User struct {
    Name           string `json:"name"`
    HashedPassword string `json:"hashedPassword"`
    Role           string `json:"role"`
    Balance        int    `json:"balance"`
}

type CreateUser struct {
    Name     string `json:"name"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

type ErrorMessage struct {
    Message string `json:"message"`
}
