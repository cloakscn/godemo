package initialize

import (
	my_router "example.com/ms_api/user_web/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	
	api_group := router.Group("/v1")
	my_router.InitUserRouter(api_group)

	return router
}