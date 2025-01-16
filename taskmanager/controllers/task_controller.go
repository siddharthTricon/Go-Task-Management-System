package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siddharthTricon/go-task-management-sysytem/repositories"
)

type TaskController struct{
	repo repositories.TaskRepository
}

func NewTaskController(repo repositories.TaskRepository) *TaskController{
	return &TaskController{repo:repo}
}  