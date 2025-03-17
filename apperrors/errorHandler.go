package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// エラーが発生したときのレスポンス処理をここで一括で行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// エラーの種類を判別して、適切な http レスポンスを返す

	// 変換先である MyAppError 型の変数を先に用意
	var appErr *MyAppError

	// errors.As 関数は、
	// - 第一引数で与えられたエラー err が、第二引数で与えられた target に変換可能であれば、結果を target 変数に格納した上で戻り値 true を返す
	// - 第一引数で与えられたエラー err が、第二引数で与えられた target に変換不可能であれば、戻り値 false を返す
	if !errors.As(err, &appErr) {
		// もし変換に失敗したら Unknown エラーを変数 appErr に手動で格納
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}
	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
