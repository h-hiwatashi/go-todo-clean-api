package controllers

import (
	"net/http"
)

// TodoControllerInterface はTodoコントローラーのインターフェースです
type TodoControllerInterface interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	GetTodo(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

// CommentControllerInterface はCommentコントローラーのインターフェースです
type CommentControllerInterface interface {
	GetComments(w http.ResponseWriter, r *http.Request)
	GetComment(w http.ResponseWriter, r *http.Request)
	CreateComment(w http.ResponseWriter, r *http.Request)
	UpdateComment(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)
}

// AppController は全てのコントローラーを結合するインターフェースです
type AppController interface {
	Todo() TodoControllerInterface
	Comment() CommentControllerInterface
}

// 1. コントローラー構造体を定義
type MyAppController struct {
	todoController    *TodoController
	commentController *CommentController
}

// AppControllerインターフェースの実装
func (c *MyAppController) Todo() TodoControllerInterface {
	return c.todoController
}

func (c *MyAppController) Comment() CommentControllerInterface {
	return c.commentController
}

// コンストラクタの定義
func NewMyAppController(todoController *TodoController, commentController *CommentController) AppController {
	return &MyAppController{
		todoController:    todoController,
		commentController: commentController,
	}
}
