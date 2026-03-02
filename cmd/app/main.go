package main

import (
	// "database/sql"
	"context"
	"log"
	"net/http"

	"task-service/internal/handler"
	"task-service/internal/repository"
	"task-service/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connString := "postgres://postgres:postgres@localhost:5432/tasks_db"
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("Error with connection:", err)
	}
	defer pool.Close()

	// repo := repository.NewTaskRepo()
	repo := repository.NewPgTaskRepo(pool)
	svc := service.NewTaskService(repo)
	h := handler.NewHandler(svc)
	
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /tasks", h.CreateTask)
	mux.HandleFunc("GET /tasks", h.GetAllTasks)
	mux.HandleFunc("GET /tasks/{id}", h.GetTask)
	mux.HandleFunc("PUT /tasks/{id}", h.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", h.DeleteTask)

	http.ListenAndServe(":8080", mux)
}
