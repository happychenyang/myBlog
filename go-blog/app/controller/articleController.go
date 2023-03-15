package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/app/service"
	. "go-blog/common/global"
	"go-blog/router"
	"net/http"
)

type Article struct {
	BaseType
}

func init() {
	userRoute := &Article{}
	//userRoute.Middleware = partial.AccountVerify()
	router.Register(userRoute, userRoute.Middleware)
}

// ArticleList 获取文章列表
func (*Article) ArticleList(c *gin.Context) {
	result, err := service.GetArticleList()
	if err != nil {
		Response(c).Error(http.StatusNotFound, err)
		return
	}
	Response(c).Success(result)
}
