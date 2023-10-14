package model

import "time"

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

type Zone struct {
	ID           int    `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	CurrentCount int    `json:"currentCount,omitempty"`
	MaxCount     int    `json:"maxCount,omitempty"`
}

type Event struct {
	ID          int       `json:"ID,omitempty"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}
