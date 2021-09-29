package env

import (
	"final_project/models"
	"github.com/apex/log"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Config models.ServerConfig //variable global

//diakses pertama kali atau dibaca pertama kali
func init() {
	err := loadConfig()
	if err != nil {
		log.Fatalf("cannot open file .env")
	}
}

func loadConfig() error {
	err := godotenv.Load() //library untuk manggil data dari env
	if err != nil {
		logrus.Fatal(err, "config/env: load file env")
	}

	// ini lib dari "github.com/caarlos0/env"
	err = env.Parse(&Config) //parsing isi dari struct server config
	if err != nil {
		return err
	}

	return err
}