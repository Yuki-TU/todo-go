package main

import (
	"log"

	"github.com/Yuki-TU/todo-go/config"
	"github.com/Yuki-TU/todo-go/routers"
	"github.com/Yuki-TU/todo-go/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// データベース初期化
	db := config.Connection()
	defer db.Close()

	if utils.GodotEnv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if utils.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	// ミドルウェアの設定
	router.Use(cors.New(cors.Config{
		// TODO: 一旦全てのドメインを許可
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	routers.SetRouting(db, router)
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}
