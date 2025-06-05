package config

import "gorm.io/gorm"

var (
	GlobalMysql Config
	DB          *gorm.DB
)
