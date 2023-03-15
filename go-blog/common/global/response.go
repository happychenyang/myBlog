package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseType struct {
	Code    int
	Message string
	Data    interface{}
	Error   interface{}
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
