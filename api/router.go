package api

import (
	"net/http"

	"github.com/h-hiwatashi/go-todo-clean-api/controllers"
	"github.com/h-hiwatashi/go-todo-clean-api/openapi"
	"github.com/h-hiwatashi/go-todo-clean-api/services"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	// 2. gorm.DB 型をもとに、サーバー全体で使用するサービス構造体 MyAppService を一つ生成する
	ser := services.NewMyAppService(db)
	// 3. MyAppService 型をもとに、サーバー全体で使用するコントローラ構造体 MyAppControllerを一つ生成する
	aCon := controllers.NewTodoController(ser)
	cCon := controllers.NewCommentController(ser)

	// CombinedController を生成
	con := controllers.NewMyAppController(aCon, cCon)

	r := openapi.Handler(con)

	return r
}
