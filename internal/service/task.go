package service

import (
	"fmt"
	"task-service/internal/model"
	"task-service/internal/repository"
	"time"
)

type TaskService interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTask(id int) (*model.Task, error)
	GetAllTasks() ([]*model.Task, error)
	UpdateTask(task *model.Task) (*model.Task, error)
	DeleteTask(id int) error
}

type MyService struct {
	repo repository.TaskRepository
}

func (s *MyService) CreateTask(task *model.Task) (*model.Task, error) {

	task.CreatedAt = time.Now()
	task.UpdatedAt = task.CreatedAt
	task.Completed = false

	err := s.repo.Create(task)
	if err != nil {
		return &model.Task{}, fmt.Errorf("save in repo error: %w", err)
	}

	return task, nil
}

func (s *MyService) GetTask(id int) (*model.Task, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error from repo: %w", err)
	}
	return task, nil 	
}

func (s *MyService) GetAllTasks() ([]*model.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("can't get tasks: %w", err)
	}
	return tasks, nil
}

func (s *MyService) UpdateTask(task *model.Task) (*model.Task, error) {
	oldTask, err := s.repo.GetByID(task.ID)
	if err != nil {
		return nil, fmt.Errorf("error from repo: %w", err)
	}
	task.CreatedAt = oldTask.CreatedAt
	task.UpdatedAt = time.Now()
	err = s.repo.Update(task)
	if err != nil {
		return nil, fmt.Errorf("error from repo: %w", err)
	}
	return task, nil
}

func (s *MyService) DeleteTask(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error from repo: %w", err)
	}
	return nil
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &MyService{repo: repo}
}
