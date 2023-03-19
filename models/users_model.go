package models

type User struct {
	UserID      uint   `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MobilePhone string `json:"mobile_phone"`
}
