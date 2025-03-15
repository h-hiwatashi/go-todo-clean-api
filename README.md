# go-todo-clean-api

TODO List API

## API endpoint

| エンドポイント    | 概要                       | レスポンスに含めたい内容     |
| ----------------- | -------------------------- | ---------------------------- |
| POST /todo        | todo を投稿する            | 投稿に成功した todo の内容   |
| GET /todo/list    | todo 一覧を取得する        | todo のリスト                |
| GET /todo/{id}    | 指定 ID の todo を取得する | todo の内容                  |
| DELETE /todo/{id} | 指定 ID の todo を削除する | todo の内容                  |
| PATCH /todo/{id}  | 指定 ID の todo を更新する | todo の内容                  |
| POST /todo/nice   | todo にいいねをつける      | いいねをつけた todo の内容   |
| POST /comment     | コメントを投稿する         | 投稿に成功したコメントの内容 |
