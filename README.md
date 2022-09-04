# todo-go
todo appのREST API

利用ライブラリ
- [gin](https://pkg.go.dev/github.com/gin-gonic/gin)
- [golang-migration](https://github.com/golang-migrate/migrate)
- [swagger-ui](https://swagger.io/tools/swagger-ui/)
- [sql](https://pkg.go.dev/database/sql)

# 環境構築

```sh
$ git clone https://github.com/Yuki-TU/todo-go
$ cd ./todo-go
$ cp .env.example .env
$ docker-compose up -d --build
$ docker-compose exec app sh
# 以下コンテナ内
$ go run main.go
```

API
[http://localhost:8081](http://localhost:8081)

swagger-ui
[http://localhost](http://localhost)


# DB

テスト環境での情報は以下

|項目|設定値|
|---|---|
|データベース種類|MySQL|
|サーバ|db|
|ユーザ|test|
|パスワード|test|
|データベース|todo|

# migration

```sh
# 環境変数定義（コンテナ起動時初回のみ）
$ export DATABASE_URI='mysql://test:test@tcp(db:3306)/todo?charset=utf8&parseTime=true&loc=Asia%2FTokyo'
# テンプレート作成
$ migrate create -ext sql -dir db/migrations -seq (table-name)
# マイグレーション適用
$ migrate -database ${DATABASE_URI} -path db/migrations up
# リバースマイグレーション
$ migrate -database ${DATABASE_URI} -path db/migrations down
```
