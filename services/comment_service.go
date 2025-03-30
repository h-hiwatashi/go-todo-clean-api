package services

import (
	"github.com/h-hiwatashi/go-todo-clean-api/apperrors"
	"github.com/h-hiwatashi/go-todo-clean-api/models"
	"github.com/h-hiwatashi/go-todo-clean-api/repositories"
)

// PostCommentHandler で使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
