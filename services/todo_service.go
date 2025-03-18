package services

import (
	"database/sql"
	"errors"

	"github.com/h-hiwatashi/go-todo-clean-api/apperrors"
	"github.com/h-hiwatashi/go-todo-clean-api/models"
	"github.com/h-hiwatashi/go-todo-clean-api/repositories"
)

// 指定IDの記事をDBから取得する関数
func (s *MyAppService) GetTodoService(todoID int) (models.Todo, error) {
	todo, err := repositories.SelectTodoDetail(s.db, todoID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Todo{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Todo{}, err
	}

	// SelectTodoDetail 関数では「指定 ID 記事に紐づいたコメント一覧」までは取得できないため、
	// SelectCommentList 関数を実行する
	commentList, err := repositories.SelectCommentList(s.db, todoID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Todo{}, err
	}

	todo.CommentList = append(todo.CommentList, commentList...)
	return todo, nil
}

// PostTodoHandler で使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func (s *MyAppService) PostTodoService(todo models.Todo) (models.Todo, error) {
	newTodo, err := repositories.InsertTodo(s.db, todo)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Todo{}, err
	}

	return newTodo, nil
}

// // TodoListHandler で使うことを想定したサービス
// // 指定 page の記事一覧を返却
// func (s *MyAppService) GetTodoListService(page int) ([]models.Todo, error) {
// 	todoList, err := repositories.SelectTodoList(s.db, page)
// 	if err != nil {
// 		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
// 		return nil, err
// 	}

// 	if len(todoList) == 0 {
// 		err := apperrors.NAData.Wrap(ErrNoData, "no data")
// 		return nil, err
// 	}

// 	return todoList, nil
// }

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func (s *MyAppService) PostNiceService(todo models.Todo) (models.Todo, error) {
	todoID := todo.ID
	err := repositories.UpdateNiceNum(s.db, int(todoID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target data")
			return models.Todo{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Todo{}, err
	}
	newTodo, err := repositories.SelectTodoDetail(s.db, int(todoID))
	if err != nil {
		return models.Todo{}, err
	}
	return newTodo, nil
}
