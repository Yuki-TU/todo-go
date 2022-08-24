package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// 指定した値の環境変数を取得
//
// @params key 取得したい環境変数キー
//
// @preturn 環境変数の値
func GodotEnv(key string) string {
	env := make(chan string, 1)

	if os.Getenv("GO_ENV") == "production" {
		env <- os.Getenv(key)
	} else {
		godotenv.Load(".env")
		env <- os.Getenv(key)
	}

	return <-env
}
