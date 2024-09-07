package db

import (
	"fmt"
	"glamurizebet/internal/application"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Create(env *application.Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(144.22.223.201:1432)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DbUser, env.DbPass, env.DbName)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
