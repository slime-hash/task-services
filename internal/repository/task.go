package repository

import (
	"errors"
	"task-service/internal/model"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id int) (*model.Task, error)
	GetAll() ([]*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}

var ErrNotFound = errors.New("not found")

type InMemoryRepo struct {
	tasks []*model.Task
}

func (r *InMemoryRepo) Create(task *model.Task) error {
	id := len(r.tasks) + 1
	task.ID = id
	r.tasks = append(r.tasks, task)
	return nil
}

func (r *InMemoryRepo) GetByID(id int) (*model.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, ErrNotFound
}

func (r *InMemoryRepo) GetAll() ([]*model.Task, error) {
	return r.tasks, nil
}

func (r *InMemoryRepo) Update(task *model.Task) error {
	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = task
			return nil
		}
	}
	return ErrNotFound
}

func (r *InMemoryRepo) Delete(id int) error {
	for i, t:= range r.tasks {
		if t.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func NewTaskRepo() TaskRepository {
	return &InMemoryRepo{tasks: make([]*model.Task, 0)}
}
