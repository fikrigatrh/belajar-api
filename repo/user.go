package repo

import (
	"final_project/models"
	"gorm.io/gorm"
)

// ini adalah tempat untuk komunikasi dengan db seperti membuat query

// 1
type UserRepoStruct struct {
	db *gorm.DB
}

//4
func NewUserRepoImpl(db *gorm.DB) UserRepoInterface {
	return &UserRepoStruct{db: db}
}

// 2
// AddData store to database
func (u *UserRepoStruct) AddData() models.User {
	usr := models.User{
		Nama:   "gatra",
		Alamat: "indonesia",
	}

	// dokumentasi gorm
	err := u.db.Create(&usr).Error
	if err != nil {
		return models.User{}
	}

	return usr
}
