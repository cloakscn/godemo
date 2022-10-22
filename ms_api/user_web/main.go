package main

import (
	"fmt"

	"example.com/ms_api/user_web/initialize"
	"go.uber.org/zap"
)


func main() {
	port := 6789

	// initailize zap logger
	initialize.Zap()
	// initialize router
	Router := initialize.Routers()

	zap.S().Infof("服务器启动，端口：%d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}