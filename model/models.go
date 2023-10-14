package model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Balance int    `json:"balance"`
	Email   string `json:"email"`
}

type CreateUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type BookZone struct {
	UserID int `json:"userID"`
	ZoneID int `json:"zoneID"`
}

type Message struct {
	Message string `json:"message"`
}

type UserStat struct {
	CoffeeCups  int `json:"coffeeCups,omitempty"`
	CompanyDays int `json:"companyDays,omitempty"`
	OfficeHours int `json:"officeHours,omitempty"`
}
