package crud

import (
	"errors"

	"todo-app-go.com/v1/src/models"
)

var ErrTodoNotFound = errors.New("todo not found in collection")

type TodoManager struct {
	Todos []models.Todo
}

func (t *TodoManager) GetAllTodos() []models.Todo {
	return t.Todos
}

func (t *TodoManager) GetById(id int) (*models.Todo, error) {
	for _, todo := range t.Todos {
		if todo.Id == id {
			return &todo, nil
		}
	}
	return nil, ErrTodoNotFound
}
