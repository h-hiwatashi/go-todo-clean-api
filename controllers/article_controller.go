package controllers

import (
	"encoding/json"
	"strconv"

	"io"
	"net/http"

	"github.com/gorilla/mux"
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
func (c *TodoController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	// if req.Method == http.MethodGet {
	// 	io.WriteString(w, "Hello, world!\n")
	// } else {
	// もし、req の中の Method フィールドが GET でなかったら
	// Invalid method というレスポンスを、405 番のステータスコードと共に返す
	// 	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	// }
	io.WriteString(w, "Hello, world!\n")
}

// GET /todo/list
// • x が数字だった場合は、記事一覧ページの x ページ目に表示されるデータを返す
// • page に対応する値が複数個送られてきた場合には、最初の値を使用する
// • x が数字ではなかった場合には、リクエストについていたパラメータの値が悪いということなので 400 番エラーを返す
// • クエリパラメータが URL についていなかった場合には、パラメータ page=1 がついていたときと同じ処理をする
// func (c *TodoController) TodoListHandler(w http.ResponseWriter, req *http.Request) {
// 	queryMap := req.URL.Query()

// 	// クエリパラメータpageを取得
// 	var page int
// 	/// もし map 型の変数 queryMap が文字列"page"をキーに持っているのであれば、p には pageキーに対応する値 queryMap["page"] が、ok には true が格納される
// 	/// もし map 型の変数 queryMap が文字列"page"をキーに持っていないのであれば、ok にはfalse が格納される
// 	if p, ok := queryMap["page"]; ok && len(p) > 0 {
// 		var err error
// 		page, err = strconv.Atoi(p[0])
// 		if err != nil {
// 			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
// 			apperrors.ErrorHandler(w, req, err)
// 			return
// 		}
// 	} else {
// 		page = 1
// 	}

// 	todoList, err := c.service.GetTodoListService(page)
// 	if err != nil {
// 		apperrors.ErrorHandler(w, req, err)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(todoList)
// }

// POST /todo のハンドラ
func (c *TodoController) PostTodoHandler(w http.ResponseWriter, req *http.Request) {
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
func (c *TodoController) TodoDetailHandler(w http.ResponseWriter, req *http.Request) {
	// strconv.Atoi で文字列を数値に変換
	todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		// 400 番エラー (BadRequest) を返す
		apperrors.ErrorHandler(w, req, err)
		return
	}

	todo, err := c.service.GetTodoService(todoID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

// DELETE /todo/:id
func (c *TodoController) DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	// strconv.Atoi で文字列を数値に変換
	todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		// 400 番エラー (BadRequest) を返す
		apperrors.ErrorHandler(w, req, err)
		return
	}

	err = c.service.DeleteTodoService(todoID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	// 204 No Content は、リクエストが成功したが、返すデータがない場合に使うステータスコード
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("success")
}

// POST /todo/nice の挙動確認用
func (c *TodoController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
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
