package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/welcome", welcome)
	router.POST("/form", form)

	router.Run(":8080")
}

func welcome(ctx *gin.Context) {
	s := ctx.DefaultQuery("firstName", "wel")
	s2 := ctx.DefaultQuery("lastName", "come")
	ctx.JSON(http.StatusOK, gin.H{
		"firstName": s,
		"lastName": s2,
	})
}

func form(ctx *gin.Context) {
	form := ctx.PostForm("form")
	ctx.JSON(http.StatusOK, gin.H{
		"form": form,
	})
}
