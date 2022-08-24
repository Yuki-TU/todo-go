# todo-go
todo appのREST API

# 環境構築

```sh
$ git clone https://github.com/Yuki-TU/todo-go
$ cd ./todo-go
$ docker-compose up -d --build
$ docker-compose exec go sh
# 以下コンテナ内
$ go run server.go
```

[http://localhost:8081/](http://localhost:8081/)にアクセス
