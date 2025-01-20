package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/siddharthTricon/go-task-management-sysytem/controllers"
)

func AuthRoutes(router *gin.Engine, authController *controllers.AuthController){
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
}