package global

import (
	"example.com/ms_api/user_web/config"
	ut "github.com/go-playground/universal-translator"
)

var (
	Trans ut.Translator
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)