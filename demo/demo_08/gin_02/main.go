package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	goodsGroup := router.Group("/good")
	{
		goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/1", goodsDetail)
		goodsGroup.GET("/add", goodsAdd)
	}

	router.Run(":8080")
}

func goodsList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "goodsList",
	})
}

func goodsDetail(ctx *gin.Context) {

}

func goodsAdd(ctx *gin.Context) {

}