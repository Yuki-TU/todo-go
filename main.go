package main

import (
	"log"

	"github.com/Yuki-TU/todo-go/routers"
	"github.com/Yuki-TU/todo-go/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.SetRouting(router)
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}
