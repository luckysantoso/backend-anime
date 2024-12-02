package configs

import (
	"gin-gorm/configs/app_config"
	"gin-gorm/configs/db_config"
)

func InitConfigs() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}