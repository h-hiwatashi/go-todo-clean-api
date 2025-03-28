package controllers

import (
	"encoding/json"

	"io"
	"net/http"

	"github.com/h-hiwatashi/go-todo-clean-api/api"
	"github.com/h-hiwatashi/go-todo-clean-api/apperrors"
	"github.com/h-hiwatashi/go-todo-clean-api/models"

	"github.com/h-hiwatashi/go-todo-clean-api/controllers/services"
)

// 1. コントローラ構造体を定義
type TodoController struct {
	// 2. フィールドに MyAppService 構造体を含める
	service services.TodoServicer
}

// コンストラクタの定義
func NewTodoController(s services.TodoServicer) *TodoController {
	return &TodoController{service: s}
}

// 他のパッケージからも参照可能な関数・変数・定数を作成するためには、その名前を大文字から始める必要があります
func (c *TodoController) GetHello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// GET /todo/list
// • x が数字だった場合は、記事一覧ページの x ページ目に表示されるデータを返す
// • page に対応する値が複数個送られてきた場合には、最初の値を使用する
// • x が数字ではなかった場合には、リクエストについていたパラメータの値が悪いということなので 400 番エラーを返す
// • クエリパラメータが URL についていなかった場合には、パラメータ page=1 がついていたときと同じ処理をする
func (c *TodoController) GetTodoList(w http.ResponseWriter, req *http.Request, params api.GetTodoListParams) {
	// queryMap := req.URL.Query()

	// // クエリパラメータpageを取得
	// var page int
	// /// もし map 型の変数 queryMap が文字列"page"をキーに持っているのであれば、p には pageキーに対応する値 queryMap["page"] が、ok には true が格納される
	// /// もし map 型の変数 queryMap が文字列"page"をキーに持っていないのであれば、ok にはfalse が格納される
	// if p, ok := queryMap["page"]; ok && len(p) > 0 {
	// 	var err error
	// 	page, err = strconv.Atoi(p[0])
	// 	if err != nil {
	// 		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
	// 		apperrors.ErrorHandler(w, req, err)
	// 		return
	// 	}
	// } else {
	// 	page = 1
	// }

	todoList, err := c.service.GetTodoListService(*params.Page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(todoList)
}

// POST /todo のハンドラ
func (c *TodoController) CreateTodo(w http.ResponseWriter, req *http.Request) {
	var reqTodo models.Todo
	println(req.Body)
	if err := json.NewDecoder(req.Body).Decode(&reqTodo); err != nil {
		// json デコードしたときに得られたエラーをラップして、エラーメッセージを出力
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	todo, err := c.service.PostTodoService(reqTodo)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// GET /todo/:id
func (c *TodoController) GetTodoById(w http.ResponseWriter, req *http.Request, id int) {
	todo, err := c.service.GetTodoService(id)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

// DELETE /todo/:id
func (c *TodoController) DeleteTodo(w http.ResponseWriter, req *http.Request, id int) {
	err := c.service.DeleteTodoService(id)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	// 204 No Content は、リクエストが成功したが、返すデータがない場合に使うステータスコード
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("success")
}

// POST /todo/nice
func (c *TodoController) IncrementNice(w http.ResponseWriter, req *http.Request) {
	var reqTodo models.Todo
	if err := json.NewDecoder(req.Body).Decode(&reqTodo); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	todo, err := c.service.PostNiceService(reqTodo)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}
