package services

import "gorm.io/gorm"

// 1. サービス構造体を定義
type MyAppService struct {

	// 2. フィールドに gorm.DB 型を含める
	db *gorm.DB
}

// コンストラクタの定義
func NewMyAppService(db *gorm.DB) *MyAppService {
	return &MyAppService{db: db}
}
