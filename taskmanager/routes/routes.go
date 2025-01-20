package routes 

import(
	"github.com/gin-gonic/gin"
	"github.com/siddharthTricon/go-task-management-sysytem/controllers"
)

func InitRoutes(router *gin.Engine, authController *controllers.AuthController, taskController *controllers.TaskController){
	AuthRoutes(router, authController)
	TaskRoutes(router, taskController)
}