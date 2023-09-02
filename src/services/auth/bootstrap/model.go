package bootstrap

import (
	"auth/models"
	"github.com/jinzhu/gorm"
)

func Models(db *gorm.DB) {
	DB := db
	var user models.User
	DB.AutoMigrate(&user)
}
