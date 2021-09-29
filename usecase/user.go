package usecase
// folder usecase adalah tempat untuk bisnis logic / logic

import (
	"final_project/models"
	"final_project/repo"
)

// 1
type UserUsecaseStruct struct {
	usr repo.UserRepoInterface
}

func NewUseUsecaseImpl(usr repo.UserRepoInterface) UserUsecaseInterface {
	return &UserUsecaseStruct{usr: usr}
}

//2
func (u UserUsecaseStruct) Hmmm() models.User {
	usr := u.usr.AddData()

	return usr
}
