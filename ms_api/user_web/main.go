package main

import (
	"fmt"

	"example.com/ms_api/user_web/global"
	"example.com/ms_api/user_web/initialize"
	"go.uber.org/zap"
)


func main() {

	// initialize zap logger
	initialize.Zap()
	// initialize config file
	initialize.Config()
	// initialize router
	Router := initialize.Routers()
	// initialize translator
	if err := initialize.Trans("zh"); err != nil {
		zap.S().Panicln(err)
	}

	zap.S().Infof("服务器启动，端口：%s", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%s", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}