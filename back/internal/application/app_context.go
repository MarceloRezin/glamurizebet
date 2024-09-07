package application

import "gorm.io/gorm"

type AppContext struct {
	DB *gorm.DB
}
