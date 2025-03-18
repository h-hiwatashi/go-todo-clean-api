package controllers

import (
	"encoding/json"
	// "errors"

	"net/http"

	"github.com/h-hiwatashi/go-todo-clean-api/apperrors"
	"github.com/h-hiwatashi/go-todo-clean-api/models"

	"github.com/h-hiwatashi/go-todo-clean-api/controllers/services"
)

// 1. コントローラ構造体を定義
type CommentController struct {
	// 2. フィールドに MyAppService 構造体を含める
	service services.CommentServicer
}

// コンストラクタの定義
func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
