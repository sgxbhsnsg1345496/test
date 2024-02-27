package global

import (
	"gorm.io/gorm"
	"mxshop/user/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)
