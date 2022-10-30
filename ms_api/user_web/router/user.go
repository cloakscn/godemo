package router

import (
	"example.com/ms_api/user_web/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	{
		userRouter.GET("list", api.GetUserList)
		userRouter.POST("pwd_login", api.PasswordLogin)
	}
}
