package repositories

import (
	// "github.com/go-playground/validator/v10/translations/id"
	"github.com/siddharthTricon/go-task-management-sysytem/database"
	"github.com/siddharthTricon/go-task-management-sysytem/models"
)

type TaskRepository interface{
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (models.Task, error)
	CreateTask(task models.Task) error
	UpdateTask(task models.Task) error
	DeleteTask(id uint) error
}

type TaskRepositroryImpl struct{}

func NewTaskRepository() *TaskRepositroryImpl{
	return &TaskRepositroryImpl{}
}

func (repo *TaskRepositroryImpl) GetAllTasks() ([]models.Task, error){
	var tasks []models.Task
	result := database.DB.Find(&tasks)
	return tasks, result.Error 
}

func (repo *TaskRepositroryImpl) GetTaskByID(id uint) (models.Task, error){
	var task models.Task
	result := database.DB.First(&task, id)
	return task, result.Error
}

func (repo *TaskRepositroryImpl) CreateTask(task models.Task) error{
	return database.DB.Create(&task).Error
}

func (repo *TaskRepositroryImpl) UpdateTask(task models.Task) error{
	return database.DB.Save(&task).Error 
}

func (repo *TaskRepositroryImpl) DeleteTask(id uint) error{
	return database.DB.Delete(&models.Task{}, id).Error
}