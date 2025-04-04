package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"` // エラーチェーンのための内部エラー
	// Err フィールドに json:"-"というタグを付与することで、そのフィール
	// ドは json エンコードされないようにすることができます
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
