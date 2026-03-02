package repository

import (
	"task-service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgTaskRepo struct {
	db *pgxpool.Pool
}

func (r *PgTaskRepo) Create(task *model.Task) error {
	return nil
}

func (r *PgTaskRepo) GetByID(id int) (*model.Task, error) {
	return nil, nil
}

func (r *PgTaskRepo) GetAll() ([]*model.Task, error) {
	return nil, nil
}

func (r *PgTaskRepo) Update(task *model.Task) error {
	return nil
}

func (r *PgTaskRepo) Delete(id int) error {
	return nil
}

func NewPgTaskRepo(db *pgxpool.Pool) TaskRepository {
	return &PgTaskRepo{db: db}
}
