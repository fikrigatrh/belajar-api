package controller

import (
	"final_project/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	usr usecase.UserUsecaseInterface
}

func CreateUserController(r *gin.RouterGroup, usr usecase.UserUsecaseInterface)  {
	handler := UserController{usr: usr}

	r.POST("/user",  handler.AddData)
}

func (u UserController) AddData(c *gin.Context)  {
	 user := u.usr.Hmmm()
	c.JSON(http.StatusOK, user)
}