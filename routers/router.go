package routers

import (
	"database/sql"

	userList "github.com/Yuki-TU/todo-go/controllers/user/getList"
	loginAuth "github.com/Yuki-TU/todo-go/controllers/user/login"
	registerAuth "github.com/Yuki-TU/todo-go/controllers/user/register"
	handlerUserList "github.com/Yuki-TU/todo-go/handlers/users/list"
	handlerLogin "github.com/Yuki-TU/todo-go/handlers/users/login"
	handlerRegister "github.com/Yuki-TU/todo-go/handlers/users/register"

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

	userListRepository := userList.NewRepositoryUserList(db)
	userListService := userList.NewServiceRegister(userListRepository)
	userListHandler := handlerUserList.NewHandlerUserList(userListService)

	groupRoute := router.Group("/api/v1")
	groupRoute.POST("/users", registerHandler.RegisterHandler)
	groupRoute.GET("/users", userListHandler.GetUserListHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
}
