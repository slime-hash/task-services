package repository

import (
	"task-service/internal/model"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id int) (*model.Task, error)
	GetAll() ([]*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}
