package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/app/service"
	. "go-blog/common/global"
	"go-blog/middleware/partial"
	"go-blog/router"
	"net/http"
)

func init() {
	userRoute := &BaseInfo{}
	userRoute.Middleware = partial.AccountVerify()
	router.Register(userRoute, userRoute.Middleware)
}

type BaseInfo struct {
	BaseType
}

// FindBaseInfo 查询基础信息
func (*BaseInfo) FindBaseInfo(c *gin.Context) {
	result, err := service.GetBaseInfo()
	if err != nil {
		Response(c).Error(http.StatusNotFound, err)
		return
	}
	Response(c).Success(result)
}

// RotationList 轮播列表
func (*BaseInfo) RotationList(c *gin.Context) {
	result, err := service.GetRotationList()
	if err != nil {
		Response(c).Error(http.StatusNotFound, err)
		return
	}
	Response(c).Success(result)
}

// MenuList 菜单列表
func (*BaseInfo) MenuList(c *gin.Context) {
	result, err := service.GetMenuList()
	if err != nil {
		Response(c).Error(http.StatusNotFound, err)
		return
	}
	Response(c).Success(result)
}
