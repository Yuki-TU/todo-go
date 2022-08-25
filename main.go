package main

import (
	"log"

	"github.com/Yuki-TU/todo-go/routers"
	"github.com/Yuki-TU/todo-go/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// TODO: 一旦全てのドメインを許可
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	routers.SetRouting(router)
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}
