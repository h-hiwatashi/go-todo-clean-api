package repositories

import (
	"github.com/h-hiwatashi/go-todo-clean-api/models"
	"gorm.io/gorm"
)

// 新規投稿をデータベースに insert する関数
// -> データベースに保存したコメント内容と、発生したエラーを返り値にする
// func InsertComment(db *gorm.DB, comment models.Comment) (models.Comment, error) {
// 	const sqlStr = `
// 		insert into comments (article_id, message, created_at) values
// 		(?, ?, now());
// 	`
// 	var newComment models.Comment
// 	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

// 	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
// 	if err != nil {
// 		return models.Comment{}, err
// 	}

// 	id, _ := result.LastInsertId()
// 	newComment.CommentID = int(id)

// 	return newComment, nil
// }

// 指定 ID の記事についたコメント一覧を取得する関数
// -> 取得したコメントデータと、発生したエラーを返り値にする
func SelectCommentList(db *gorm.DB, articleID int) ([]models.Comment, error) {
	var commentList []models.Comment

	res := db.Find(&commentList)
	if err := res.Error; err != nil {
		return []models.Comment{}, err
	}
	return commentList, nil
}
