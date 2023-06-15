package models

import (
	"github.com/jinzhu/gorm"
)

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

type User_URL struct {
	gorm.Model

	Short      string
	Comments   string
	Medium     string
	Source     string
	Campaign   string
	Keyword    string
	Created_by string

	// Health     bool
}

type URL_INFO_RESPONSE struct {
	Url        string `json:"url"`
	Short      string `json:"short_url"`
	Count      int    `json:"Hits"`
	Comments   string `json:"Comments"`
	Medium     string `json:"Medium"`
	Source     string `json:"Source"`
	Campaign   string `json:"Campaign"`
	Keyword    string `json:"Keyword"`
	Created_by string `json:"Created_by,omitempty"`
	Created_at int64  `json:"Created_at"`
}
