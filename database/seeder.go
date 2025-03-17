package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/h-hiwatashi/go-todo-clean-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbUser     = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbConn     = fmt.Sprintf(
		"%s:%s@tcp(mysql:3306)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbDatabase,
	)
)

func main() {
	// データベースへの接続
	db, err := gorm.Open(mysql.Open(dbConn), &gorm.Config{})
	if err != nil {
		log.Println("fail to connect DB")
		log.Println(err.Error())
		return
	}
	// defer db.Close()

	// ダミーデータの生成と挿入
	for i := 0; i < 100; i++ {
		var todo models.Todo
		gofakeit.Struct(&todo)

		// TODO: バグがあるので修正する
		// データベースに挿入するクエリの作成と実行
		result := db.Create(&todo)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
}
