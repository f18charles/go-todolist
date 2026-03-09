package repository

import (
	"database/sql"
	"go-todolist/internal/models"
)

type SQLiteRepo struct {
	DB *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{DB: db}
}

func (r *SQLiteRepo) Create(todo models.Todo) error {
	_, err := r.DB.Exec(
		"INSERT INTO todos(id,title,completed,created_at) VALUES(?,?,?,?)",
		todo.ID, todo.Title, todo.Completed, todo.CreatedAt,
	)
	return err
}

func (r *SQLiteRepo) GetByID(id string) (models.Todo, error) {
	row := r.DB.QueryRow(
		"SELECT id,title,completed,created_at FROM todos WHERE id=?",
		id,
	)

	var t models.Todo
	err := row.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
	return t, err
}

func (r *SQLiteRepo) GetAll() ([]models.Todo, error) {
	rows, err := r.DB.Query(
		"SELECT id,title,completed,created_at FROM todos",
	)
	if err != nil {
		return nil, err
	}

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *SQLiteRepo) Update(todo models.Todo) error {
	_, err := r.DB.Exec(
		"UPDATE todos SET title=?,completed=? WHERE id=?",
		todo.Title, todo.Completed, todo.ID,
	)
	return err
}

func (r *SQLiteRepo) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM todos WHERE id=?", id)
	return err
}
