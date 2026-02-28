package model

import (
	"time"
)

type Task struct {
	ID 			int			`json:"id"`
	Title 		string		`json:"title"`
	Description string		`json:"description"`
	Completed	bool		`json:"completed"`
	Deadline	*time.Time	`json:"deadline"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type CreateTaskRequest struct {
	Title		string		`json:"title"`
	Description	string		`json:"description,omitempty"`
	Deadline	*time.Time	`json:"deadline,omitempty"`
}
