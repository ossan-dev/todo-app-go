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

func (t *TodoManager) GetByStatus(completedStatus bool) []models.Todo {
	result := make([]models.Todo, 0)
	for _, todo := range t.Todos {
		if todo.IsCompleted == completedStatus {
			result = append(result, todo)
		}
	}
	return result
}

func (t *TodoManager) Add(todo models.Todo) {
	// TODO: add validation logic & tests
	t.Todos = append(t.Todos, todo)
}

func (t *TodoManager) Update(todo models.Todo) error {
	for key, value := range t.Todos {
		if value.Id == todo.Id {
			t.Todos[key] = todo
			return nil
		}
	}
	return ErrTodoNotFound
}

func (t *TodoManager) DeleteById(id int) error {
	index := -1
	for key, value := range t.Todos {
		if value.Id == id {
			index = key
			return nil
		}
	}
	t.Todos = append(t.Todos[:index], t.Todos[index+1:]...)
	return nil
}
