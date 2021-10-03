package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama   string `json:"nama"`
	Alamat interface{} `json:"alamat"`
}

type AlamatDetail struct {
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
}

type UserAPI struct {
	Sammy UserAPIDetail `json:"sammy"`
	Jesse UserAPIDetail `json:"jesse"`
	Drew  UserAPIDetail `json:"drew"`
	Jamie UserAPIDetail `json:"jamie"`
 }

type UserAPIDetail struct {
	Username string `json:"username,omitempty"`
	Followers int `json:"followers,omitempty"`
}

