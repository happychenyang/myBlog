package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/app/bean"
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
	articleWhere := &bean.ArticleList{}
	if err := c.ShouldBindJSON(articleWhere); err != nil {
		Response(c).Error(http.StatusNotFound, err.Error())
		return
	}
	result, total, err := service.GetArticleList(articleWhere)
	if err != nil {
		Response(c).Error(http.StatusNotFound, err.Error())
		return
	}
	Response(c).JsonPagination(&articleWhere.Page, result, total)
}

// ArticleDetail 获取文章详情
func (*Article) ArticleDetail(c *gin.Context) {
	var articleWhere bean.ArticleDetail
	if err := c.ShouldBindJSON(&articleWhere); err != nil {
		Response(c).Error(http.StatusNotFound, err.Error())
		return
	}
	result, err := service.GetArticleDetail(articleWhere)
	if err != nil {
		Response(c).Error(http.StatusNotFound, err)
		return
	}
	Response(c).Success(result)
}

// ArticleClassify 获取文章分类
func (*Article) ArticleClassify(c *gin.Context) {
	result, err := service.GetArticleClassify()
	if err != nil {
		Response(c).Error(http.StatusNotFound, err)
		return
	}
	Response(c).Success(result)
}
