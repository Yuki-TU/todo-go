package routers

import (
	"database/sql"

	loginAuth "github.com/Yuki-TU/todo-go/controllers/user/login"
	registerAuth "github.com/Yuki-TU/todo-go/controllers/user/register"
	handlerRegister "github.com/Yuki-TU/todo-go/handlers/users"
	handlerLogin "github.com/Yuki-TU/todo-go/handlers/users/login"
	"github.com/gin-gonic/gin"
)

// ルーティングの設定を行う
//
// @param router ルーター
func SetRouting(db *sql.DB, router *gin.Engine) {
	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	loginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(loginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	groupRoute := router.Group("/api/v1")
	groupRoute.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	groupRoute.POST("/users", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
}
