package repositories

import (
	"log"

	"github.com/h-hiwatashi/go-todo-clean-api/models"
	"gorm.io/gorm"
)

const (
	todoNumPerPage = 5
)

// 新規投稿をデータベースに insert する関数
// -> データベースに保存した記事内容と、発生したエラーを返り値にする
func InsertTodo(db *gorm.DB, todo models.Todo) (models.Todo, error) {
	// const sqlStr = `
	// insert into todos (title, contents, username, nice, created_at) values
	// (?, ?, ?, 0, now());
	// `

	// var newTodo models.Todo
	// newTodo.Title, newTodo.Contents, newTodo.UserName = todo.Title, todo.Contents, todo.UserName

	// result, err := db.Exec(sqlStr, todo.Title, todo.Contents, todo.UserName)
	result := db.Create(&todo)
	if result.Error != nil {
		return models.Todo{}, result.Error
	}

	// id := todo.ID

	// newTodo.ID = int(id)

	return todo, nil
}

// 変数 page で指定されたページに表示する投稿一覧をデータベースから取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectTodoList(db *gorm.DB, page int) ([]models.Todo, error) {
	var todos []models.Todo
	offset := (page - 1) * todoNumPerPage
	// rows, err := db.Query(sqlStr, todoNumPerPage, offset)
	result := db.Limit(todoNumPerPage).Offset(offset).Find(&todos)
	if result.Error != nil {
		log.Printf("Error db.Query: %v\n", result.Error)
		return nil, result.Error
	}
	return todos, nil
}

// 投稿 ID を指定して、記事データを取得する関数
// -> 取得した記事データと、発生したエラーを返り値にする
func SelectTodoDetail(db *gorm.DB, todoID int) (models.Todo, error) {
	var todo models.Todo

	res := db.First(&todo, todoID)
	if err := res.Error; err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func DeleteTodo(db *gorm.DB, todoID int) error {
	var todo models.Todo

	res := db.Delete(&todo, todoID)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

// いいねの数を update する関数
// -> 発生したエラーを返り値にする
func UpdateNiceNum(db *gorm.DB, todoID int) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 既存のTodoをIDで取得
	var todo models.Todo
	if err := tx.First(&todo, todoID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// いいね数をインクリメント
	todo.NiceNum++

	// 更新したTodoをデータベースに保存
	if err := tx.Save(&todo).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
