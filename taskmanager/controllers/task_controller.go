package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/utils"
	"github.com/siddharthTricon/go-task-management-sysytem/models"
	"github.com/siddharthTricon/go-task-management-sysytem/repositories"
)

type TaskController struct{
	repo repositories.TaskRepository
}

func NewTaskController(repo repositories.TaskRepository) *TaskController{
	return &TaskController{repo:repo}
}  

func (tc *TaskController) GetAllTasks(c *gin.Context){
	tasks,  err := tc.repo.GetAllTasks()

	if err != nil{
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusOK, tasks)
}

func (tc *TaskController) CreateTask(c *gin.Context){
	var task models.Task

	if err := c.BindJSON(&task); err != nil{
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tc.repo.CreateTask(task); err != nil{
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusCreated, task)
}