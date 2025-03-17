package services

import "github.com/h-hiwatashi/go-todo-clean-api/models"

// /todo 関連を引き受けるサービス
type TodoServicer interface {
	PostTodoService(todo models.Todo) (models.Todo, error)
	// GetTodoListService(page int) ([]models.Todo, error)
	// GetTodoService(todoID int) (models.Todo, error)
	// PostNiceService(todo models.Todo) (models.Todo, error)
}

// /comment を引き受けるサービス
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
