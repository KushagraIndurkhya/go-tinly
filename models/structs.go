package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id    string `gorm:"PRIMARY_KEY"`
	Email string
	Name  string
	Limit int
	// URL   pq.StringArray `gorm:"type:varchar(64)[]"`
	//TODO: Add Access & Refresh token
}
type UserInfo struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type User_URL struct {
	gorm.Model

	Short      string
	Created_by string
	// Health     bool
}

type URL_INFO struct {
	Url        string    `json:"url"`
	Short      string    `json:"short_url"`
	Count      int       `json:"Hits"`
	Created_by string    `json:"Created_by,omitempty"`
	Created_at time.Time `json:"Created_at"`
}
