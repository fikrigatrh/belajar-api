package repo

import "final_project/models"

// 3
type UserRepoInterface interface {
	// AddData list all method function in package repo
	AddData(usr models.User) models.User
	ConsumeApi() models.UserAPI
}
