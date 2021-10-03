package repo

import (
	"encoding/json"
	"final_project/config/env"
	"final_project/models"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type UserRepoStruct struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) UserRepoInterface {
	return &UserRepoStruct{db: db}
}


func (u *UserRepoStruct) AddData(usr models.User) models.User {

	err := u.db.Create(&usr).Error
	if err != nil {
		return models.User{}
	}

	return usr
}

func (u *UserRepoStruct) ConsumeApi() models.UserAPI {

	client := resty.New()
	url := env.Config.UrlJsonUser

	res := models.UserAPI{}
	resp, err := client.R().
		SetResult(res).
		Get(url)

	if err != nil {
		return models.UserAPI{}
	}

	result := resp.String()
	if errs := json.Unmarshal([]byte(result), &res); errs != nil {
		return models.UserAPI{}
	}

	return res
}