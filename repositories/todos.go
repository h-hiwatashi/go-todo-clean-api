package repositories

import (
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
// func SelectTodoList(db *gorm.DB, page int) ([]models.Todo, error) {
// 	const sqlStr = `
// 		select todo_id, title, contents, username, nice
// 		from todos
// 		limit ? offset ?;
// 	`

// 	rows, err := db.Query(sqlStr, todoNumPerPage, ((page - 1) * todoNumPerPage))
// 	if err != nil {
// 		log.Printf("Error db.Query: %v\n", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	todoArray := make([]models.Todo, 0)
// 	for rows.Next() {
// 		var todo models.Todo
// 		rows.Scan(&todo.ID, &todo.Title, &todo.Contents, &todo.UserName, &todo.NiceNum)

// 		todoArray = append(todoArray, todo)
// 	}

// 	return todoArray, nil
// }

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

// // いいねの数を update する関数
// // -> 発生したエラーを返り値にする
// func UpdateNiceNum(db *gorm.DB, todoID int) error {
// 	const sqlGetNice = `
// 	select nice
// 	from todos
// 	where todo_id = ?;
// 	`
// 	const sqlUpdateNice = `update todos set nice = ? where todo_id = ?`

// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	row := tx.QueryRow(sqlGetNice, todoID)
// 	if err := row.Err(); err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	var nicenum int
// 	err = row.Scan(&nicenum)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	_, err = tx.Exec(sqlUpdateNice, nicenum+1, todoID)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	tx.Commit()
// 	return nil
// }
