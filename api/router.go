package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/h-hiwatashi/go-todo-clean-api/controllers"
	"github.com/h-hiwatashi/go-todo-clean-api/services"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *mux.Router {
	// 2. gorm.DB 型をもとに、サーバー全体で使用するサービス構造体 MyAppService を一つ生成する
	ser := services.NewMyAppService(db)
	// 3. MyAppService 型をもとに、サーバー全体で使用するコントローラ構造体 MyAppControllerを一つ生成する
	aCon := controllers.NewTodoController(ser)
	// cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo", aCon.PostTodoHandler).Methods(http.MethodPost)
	// r.HandleFunc("/todo/list", aCon.TodoListHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo/{id:[0-9]}", aCon.TodoDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/todo/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	// r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
