package testdata

import "github.com/h-hiwatashi/go-todo-clean-api/models"

var todoTestData = []models.Todo{
	models.Todo{
		ID:          1,
		Title:       "firstPost",
		Contents:    "This is my first blog",
		UserName:    "user",
		NiceNum:     2,
		CommentList: commentTestData,
	},
	models.Todo{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "user",
		NiceNum:  4,
	},
}
var commentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		TodoID:    1,
		Message:   "1st comment yeah",
	},
	models.Comment{
		CommentID: 2,
		TodoID:    1,
		Message:   "welcome",
	}}
