package service

import (
	"time"

	"github.com/google/uuid"
	"go-todolist/internal/models"
	"go-todolist/internal/repository"
)

type TodoService struct {
	Repo repository.TodoRepository
}

func NewTodoService(r repository.TodoRepository) *TodoService {
	return &TodoService{Repo: r}
}

func (s *TodoService) Create(title string) (models.Todo, error) {
	todo := models.Todo{
		ID:        uuid.New().String(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	err := s.Repo.Create(todo)

	return todo, err
}

func (s *TodoService) List() ([]models.Todo, error) {
	return s.Repo.GetAll()
}

func (s *TodoService) Complete(id string) error {
	todo, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}

	todo.Completed = true

	return s.Repo.Update(todo)
}

func (s *TodoService) Delete(id string) error {
	return s.Repo.Delete(id)
}
