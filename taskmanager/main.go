

package main

import (
	"log"

	"github.com/siddharthTricon/go-task-management-sysytem/controllers"
	"github.com/siddharthTricon/go-task-management-sysytem/database"
	"github.com/siddharthTricon/go-task-management-sysytem/models"
	"github.com/siddharthTricon/go-task-management-sysytem/repositories"
	"github.com/siddharthTricon/go-task-management-sysytem/routes"
	"github.com/siddharthTricon/go-task-management-sysytem/services"
	"github.com/siddharthTricon/go-task-management-sysytem/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables and initialize logger
	utils.LoadEnv()
	utils.InitLogger()

	// Connect to the database
	database.Connect()

	// Automatically migrate database models
	if err := database.DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatal("Failed to auto-migrate models: ", err)
	}

	// Initialize services and controllers
	jwtService := services.NewJWTService()
	authController := controllers.NewAuthController(jwtService)
	taskController := controllers.NewTaskController(repositories.NewTaskRepository())

	// Create Gin router
	router := gin.Default()

	// Initialize routes
	routes.InitRoutes(router, authController, taskController)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
