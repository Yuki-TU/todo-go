package routers

import (
	"github.com/gin-gonic/gin"
)

// ルーティングの設定を行う
//
// @param router ルーター
func SetRouting(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"email":    "hoge@hoge.com",
			"fullname": "GET",
			"passwrod": "成功しました。",
		})
	})
}
