package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Good struct {
	ID     int    `uri:"id" binding:"required"`
	Action string `uri:"action" binding:"required"`
}

func main() {
	router := gin.Default()
	goodsGroup := router.Group("/good")
	{
		goodsGroup.GET("", goodsList)
		// todo get goods id
		goodsGroup.GET("/:id/:action", goodsDetail)
		goodsGroup.POST("", goodsAdd)
	}

	router.Run(":8080")
}

func goodsList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "goodsList",
	})
}

func goodsDetail(ctx *gin.Context) {
	var good Good
	if err := ctx.ShouldBindUri(&good); err != nil {
		ctx.Status(http.StatusNotFound)
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"id":     good.ID,
			"action": good.Action,
		})
	}
}

func goodsAdd(ctx *gin.Context) {

}
