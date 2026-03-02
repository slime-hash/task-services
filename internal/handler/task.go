package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"task-service/internal/model"
	"task-service/internal/repository"
	"task-service/internal/service"
)

type TaskHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	GetTask(w http.ResponseWriter, r *http.Request)
	GetAllTasks(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type MyHandler struct {
	svc service.TaskService
}

func (h *MyHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if len(strings.TrimSpace(req.Title)) == 0 {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}
	task := &model.Task{
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
	}
	created, err := h.svc.CreateTask(task)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(created)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	task, err := h.svc.GetTask(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
	}
}

func (h *MyHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func (h *MyHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	var task model.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if len(strings.TrimSpace(task.Title)) == 0 {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}
	task.ID = id
	newTask, err := h.svc.UpdateTask(&task)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(newTask)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func (h *MyHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	err = h.svc.DeleteTask(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func NewHandler(svc service.TaskService) TaskHandler {
	return &MyHandler{svc: svc}
}
