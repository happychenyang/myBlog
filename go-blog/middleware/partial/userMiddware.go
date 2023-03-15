package partial

import (
	"github.com/gin-gonic/gin"
	model "go-blog/app/models"
	. "go-blog/common/global"
	"net/http"
)

func AccountVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.Verify
		err := ctx.ShouldBind(&user) //绑定参数
		result := user.CheckAccount(user)
		if err != nil || !result {
			Response(ctx).Error(http.StatusNotFound, "这是User中间件，账号不存在")
			ctx.Abort() //终止后续代码逻辑
		}
		ctx.Next()
	}
}
