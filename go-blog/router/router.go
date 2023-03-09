package router

import (
	"github.com/gin-gonic/gin"
	. "go-blog/app/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("./findUser", FindUser)
	return router
}
