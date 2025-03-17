package apperrors

type ErrCode string

const (
	Unknown          ErrCode = "U000" // 未知のエラー
	InsertDataFailed ErrCode = "S001" // データの挿入に失敗
	GetDataFailed    ErrCode = "S002" // データの取得に失敗
	NAData           ErrCode = "S003" // データが存在しない
	NoTargetData     ErrCode = "S004" // 対象のデータが存在しない
	UpdateDataFailed ErrCode = "S005" // データの更新に失敗

	ReqBodyDecodeFailed ErrCode = "R001" // リクエストボディのデコードに失敗
	BadParam            ErrCode = "R002" // リクエストパラメータが不正
)
