package middle_ware

import (
	"github.com/gin-gonic/gin"
	"jincheng/config"
	"jincheng/internal/core/base"
	"jincheng/internal/core/jwt"
)

func Jwt(conf *config.Config) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("auth-token")

		if context.Request.RequestURI == "/api/jc/admin/login" {
			context.Next()
			return
		}

		_, err := jwt.ParseToken(token, conf)
		if err != nil {
			base.NewResponse(context).Unauthorized("请先登录")
			context.Abort()
			return
		}
		context.Next()
	}
}
