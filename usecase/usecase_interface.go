package usecase

import "final_project/models"

type UserUsecaseInterface interface {
	Hmmm(usr models.User) models.User
	GetCountFollower(username string) models.UserAPIDetail
	GetDetail(userId string) models.UserAPIDetail
}
