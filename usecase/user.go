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
func (u UserUsecaseStruct) Hmmm(usr models.User) models.User {
	user := u.usr.AddData(usr)

	return user
}

func (u UserUsecaseStruct) GetCountFollower(username string) models.UserAPIDetail {
	var res models.UserAPIDetail
	usr := u.usr.ConsumeApi()
	switch username {
	case "SammyShark":
		res.Followers = usr.Sammy.Followers
	case "JesseOctopus":
		res.Followers = usr.Jesse.Followers
	case "DrewSquid":
		res.Followers = usr.Drew.Followers
	case "JamieMantisShrimp":
		res.Followers = usr.Jamie.Followers
	}

	return res
}

func (u UserUsecaseStruct) GetDetail(userId string) models.UserAPIDetail {
	var res models.UserAPIDetail
	usr := u.usr.ConsumeApi()
	switch userId {
	case "sammy":
		res.Username = usr.Sammy.Username
		res.Followers = usr.Sammy.Followers
	case "jesse":
		res.Username = usr.Jesse.Username
		res.Followers = usr.Jesse.Followers
	case "drew":
		res.Username = usr.Drew.Username
		res.Followers = usr.Drew.Followers
	case "jamie":
		res.Username = usr.Jamie.Username
		res.Followers = usr.Jamie.Followers
	}

	return res
}
