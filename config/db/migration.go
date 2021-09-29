package db

import (
	"final_project/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	var err error

	err = db.Debug().AutoMigrate(models.User{}, models.ServerConfig{})
	if err != nil {
		return
	}

}