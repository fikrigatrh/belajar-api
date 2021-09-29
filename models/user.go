package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

type alamatDetail struct {
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
}
