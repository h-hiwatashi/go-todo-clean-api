package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/h-hiwatashi/go-todo-clean-api/models"
)

var (
	dbUser     = os.Getenv("MYSQL_DATABASE_USER")
	dbPassword = os.Getenv("MYSQL_DATABASE_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbConn     = fmt.Sprintf(
		"%s:%s@tcp(mysql:3306)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbDatabase,
	)
)

func main() {
	// 1. サーバー全体で使用する sql.DB 型を一つ生成する。
	// db, err := sql.Open("mysql", dbConn)
	db, err := gorm.Open(mysql.Open(dbConn), &gorm.Config{})
	if err != nil {
		log.Println("fail to connect DB")
		log.Println(err.Error())
		return
	}

	// db を migrate する。
	db.AutoMigrate(&models.Todo{}, &models.Comment{})

	// 4. コントローラ型 MyAppController のハンドラメソッドとパスとの関連付けを行う。
	// r := api.NewRouter(db)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")

	// ListenAndServe 関数にて、サーバーを起動
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// log.Fatal(http.ListenAndServe(":8080", r))
}
