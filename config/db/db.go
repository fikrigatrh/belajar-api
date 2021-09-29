package db
// SETTING CONNECTION TO DB

import (
	"final_project/config/env"
	"fmt"
	log2 "github.com/apex/log"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // variable global DB

func init() {
	var err error
	var configDb = env.Config

	if DB != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log2.Error("recovery error")
		}
	}()

	urlDB := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",configDb.Host,configDb.Username,configDb.Password,configDb.Name,configDb.Port)
	DB, err = gorm.Open(postgres.Open(urlDB), &gorm.Config{})
	if err != nil {
		log2.Fatal("cannot connect to DB")
	}

	log.Print("success connect to db")

	AutoMigrate(DB)

}
