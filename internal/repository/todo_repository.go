package repository

import "go-todolist/internal/models"

type TodoRepository interface {
	Create(todo models.Todo) error
	GetByID(id string) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	Update(todo models.Todo) error
	Delete(id string) error
}
