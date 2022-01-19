package crud

import (
	"errors"

	"todo-app-go.com/v1/src/models"
)

var ErrTodoNotFound = errors.New("todo not found in collection")

type TodoStore interface {
	GetTodoById(id int) string
}

type TodoManager struct {
	Todos []models.Todo
}

func (t *TodoManager) GetAllTodos() []models.Todo {
	return t.Todos
}

func GetTodoById(id int) models.Todo {
	if id == 1 {
		return models.Todo{1, "FirstTodo", false}
	}

	if id == 2 {
		return models.Todo{2, "SecondTodo", false}
	}

	return models.Todo{}
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
		}
	}

	if index < 0 {
		return ErrTodoNotFound
	}

	t.Todos = append(t.Todos[:index], t.Todos[index+1:]...)
	return nil
}
