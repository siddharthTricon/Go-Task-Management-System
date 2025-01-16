package models

type Task struct{
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	AssignedTo uint `json:"assigned_to"`
	Status string `json:"status"`
}