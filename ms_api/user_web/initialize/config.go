package initialize

import (
	"fmt"

	"example.com/ms_api/user_web/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func Config() {
	b := GetEnvInfo("MS_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("./%s-prod.yaml", configFilePrefix)
	if b {
		configFileName = fmt.Sprintf("./%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf(err.Error())
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Panicf(err.Error())
	}

	go func() {
		v.WatchConfig()
		for {
			v.OnConfigChange(func(in fsnotify.Event) {
				zap.S().Infof("config file change: %s", in.Name)

				if err := v.ReadInConfig(); err != nil {
					zap.S().Panicf(err.Error())
				}
				if err := v.Unmarshal(global.ServerConfig); err != nil {
					zap.S().Panicf(err.Error())
				}
			})
		}
	}()
}
