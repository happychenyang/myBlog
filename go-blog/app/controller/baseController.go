package controller

import "github.com/gin-gonic/gin"

type BaseType struct {
	Middleware gin.HandlerFunc
	UserName   string
}
