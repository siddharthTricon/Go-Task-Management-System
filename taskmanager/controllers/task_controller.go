package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/siddharthTricon/go-task-management-sysytem/models"
	"github.com/siddharthTricon/go-task-management-sysytem/repositories"
	"github.com/siddharthTricon/go-task-management-sysytem/utils"
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
		utils.RespondJSON(c, http.StatusInternalServerError,"Error fetching task", gin.H{"error":err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusOK,"Task fetched successfully", tasks)
}

func (tc *TaskController) CreateTask(c *gin.Context){
	var task models.Task

	if err := c.BindJSON(&task); err != nil{
		utils.RespondJSON(c, http.StatusInternalServerError,"Invalid task data", gin.H{"error": err.Error()})
		return
	}

	if err := tc.repo.CreateTask(task); err != nil{
		utils.RespondJSON(c, http.StatusInternalServerError,"Error creating task", gin.H{"error":err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusCreated,"Task created", task)
}