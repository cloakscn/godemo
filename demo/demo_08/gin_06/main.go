package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignInForm struct {
	User     string `json:"user" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required"`
}

type SignUpForm struct {
	Age uint8 `form:"age" binding:"gte=1,lte=130"`
	Name string `form:"name" binding:"required,min=3"`
	Email string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/loginJSON", func(ctx *gin.Context) {
		var loginForm SignInForm
		err := ctx.ShouldBind(&loginForm)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	})

	router.POST("/signUp", func(ctx *gin.Context) {
		var signUpForm SignUpForm
		err := ctx.ShouldBind(&signUpForm)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})

	router.Run(":8083")
}
