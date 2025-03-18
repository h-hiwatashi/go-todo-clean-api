package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	TodoID  int    `json:"todo_id" gorm:"not null"`
	Message string `json:"message"`
}

type Todo struct {
	gorm.Model
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comments"`
}
