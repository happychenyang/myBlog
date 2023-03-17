package global

import (
	"github.com/gin-gonic/gin"
	"go-blog/app/bean"
	"net/http"
)

type ResponseType struct {
	Code      int
	Message   string
	PageIndex int
	PageSize  int
	Total     int64
	Data      interface{}
	Error     interface{}
}

type GinContext struct {
	c *gin.Context
}

// Response 使用返回类
func Response(c *gin.Context) *GinContext {
	return &GinContext{c: c}
}

// Success 成功后返回
func (GinContext *GinContext) Success(data interface{}) {
	successInfo := ResponseType{}
	successInfo.Message = "success"
	successInfo.Data = data
	GinContext.c.JSON(http.StatusOK, successInfo)
}

// Error 错误类返回
func (GinContext *GinContext) Error(httpCode int, error interface{}) {
	errorInfo := ResponseType{}
	errorInfo.Message = "error"
	errorInfo.Error = error
	GinContext.c.JSON(httpCode, errorInfo)
}

// JsonPagination json分页
func (GinContext *GinContext) JsonPagination(page *bean.Page, list interface{}, total int64) {
	successInfo := ResponseType{}
	successInfo.Message = "success"
	successInfo.Total = total
	successInfo.Data = list
	successInfo.PageSize = page.PageSize
	successInfo.PageIndex = page.PageIndex
	GinContext.c.JSON(http.StatusOK, successInfo)
}
