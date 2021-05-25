package models

type User struct {
	Id    string `gorm:"PRIMARY_KEY"`
	Email string
	Name  string
	Limit int
	//TODO: Add Access & Refresh token
}
type UserInfo struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
