package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/siddharthTricon/go-task-management-sysytem/controllers"
	"github.com/siddharthTricon/go-task-management-sysytem/middlewares"
)

func TaskRoutes(router *gin.Engine, taskController *controllers.TaskController){
	taskGroup := router.Group("/task")
	taskGroup.Use(middlewares.AuthMiddleware())
	{
		taskGroup.GET("/", taskController.GetAllTasks)
		taskGroup.POST("/", taskController.CreateTask)
	// 	taskGroup.GET("/:id", taskController.GetTaskByID)
	// 	taskGroup.PUT("/:id", taskController.UpdateTask)
	// 	taskGroup.DELETE("/:id", taskController.DeleteTask)
	}
}