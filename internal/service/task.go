package service

import (
	"task-service/internal/model"
)

type TaskService interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTask(id int) (*model.Task, error)
	GetAllTasks() ([]*model.Task, error)
	UpdateTask(task *model.Task) (*model.Task, error)
	DeleteTask(id int) error
}
