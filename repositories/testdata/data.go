package testdata

import (
	"github.com/h-hiwatashi/go-todo-clean-api/models"
)

var TodoTestData = []models.Todo{
	models.Todo{
		ID:       1,
		Title:    "first Title",
		Contents: "This is my first todo",
		UserName: "user name",
		NiceNum:  2,
	},
	models.Todo{
		ID:       2,
		Title:    "2nd",
		Contents: "Second todo post",
		UserName: "user name",
		NiceNum:  4,
	}}
