package bootstrap

import (
	"gateway/models"
	"github.com/jinzhu/gorm"
)

func Models(db *gorm.DB) {
	DB := db
	var user models.User
	DB.AutoMigrate(&user)
}
