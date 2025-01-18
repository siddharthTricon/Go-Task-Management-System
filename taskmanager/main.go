// package main

// import (
// "github.com/siddharthTricon/go-task-management-sysytem/controllers"
// "github.com/siddharthTricon/go-task-management-sysytem/database"
// "github.com/siddharthTricon/go-task-management-sysytem/middlewares"
// "github.com/siddharthTricon/go-task-management-sysytem/repositories"
// "github.com/siddharthTricon/go-task-management-sysytem/services"
// "github.com/siddharthTricon/go-task-management-sysytem/utils"
//     "github.com/gin-gonic/gin"
// )

// func main() {
//     utils.LoadEnv()
//     utils.InitLogger()

//     database.Connect()
//     database.DB.AutoMigrate(&models.User{}, &models.Task{})

//     jwtService := services.NewJWTService()
//     authController := controllers.NewAuthController(jwtService)
//     taskController := controllers.NewTaskController(repositories.NewTaskRepository())

//     router := gin.Default()

//     // Auth Routes
//     router.POST("/register", authController.Register)
//     router.POST("/login", authController.Login)

//     // Task Routes (protected)
//     taskRoutes := router.Group("/tasks")
//     taskRoutes.Use(middlewares.AuthMiddleware())
//     {
//         taskRoutes.GET("/", taskController.GetAllTasks)
//         taskRoutes.POST("/", taskController.CreateTask)
//     }

//     router.Run(":8080")
// }








package main

import (
	"taskmanager/controllers"
	"taskmanager/database"
	"taskmanager/repositories"
	"taskmanager/routes"
	"taskmanager/services"
	"taskmanager/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	utils.InitLogger()

	// Database connection and migrations
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Task{})

	// Initialize services and controllers
	jwtService := services.NewJWTService()
	authController := controllers.NewAuthController(jwtService)
	taskController := controllers.NewTaskController(repositories.NewTaskRepository())

	// Create router
	router := gin.Default()

	// Register routes
	routes.InitRoutes(router, authController, taskController)

	// Start server
	router.Run(":8080")
}
