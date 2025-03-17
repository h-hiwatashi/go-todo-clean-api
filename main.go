package main

import (
	"fmt"
	"log"
	"os"
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
	// 1. サーバー全体で使用する sql.DB 型を一つ生成する
	// db, err := sql.Open("mysql", dbConn)
	// if err != nil {
	// 	log.Println("fail to connect DB")
	// 	return
	// }
	// 4. コントローラ型 MyAppController のハンドラメソッドとパスとの関連付けを行う
	// r := api.NewRouter(db)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")
}
