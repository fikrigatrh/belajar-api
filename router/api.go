package router

import (
	"final_project/config/db"
	"final_project/controller"
	"final_project/repo"
	"final_project/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.New() // CALL LIBRARY GIN GONIC FOR ROUTER
	DB := db.DB // CALL FUNCTION DB
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CALL DEPENDENCY REPOSITORY
	userRepo := repo.NewUserRepoImpl(DB)

	// CALL DEPENDENCY USECASE
	userUsecase := usecase.NewUseUsecaseImpl(userRepo)

	newRoute := router.Group("/paymentMonitoring")
	// CALL DEPENDENCY CONTROLLER
	controller.CreateUserController(newRoute, userUsecase)

	// RETURN ROUTER
	return router
}
