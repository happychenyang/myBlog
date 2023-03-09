package controller

import (
	"github.com/gin-gonic/gin"
	model "go-blog/models"
	"net/http"
)

func FindUser(c *gin.Context) {
	var user model.User
	result, err := user.FindUserList()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "noe found",
			"data": "",
			"err":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": result,
		"err":  "",
	})
}
