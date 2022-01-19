package stubs

import (
	"errors"

	"todo-app-go.com/v1/src/models"
)

type StubTodoStore struct {
	todos map[int]models.Todo
}

func NewStubTodoStore(todos *map[int]models.Todo) *StubTodoStore {
	return &StubTodoStore{todos: *todos}
}

func (s *StubTodoStore) GetTodoById(id int) (string, error) {
	if val, ok := s.todos[id]; ok {
		todo := val
		return todo.Description, nil
	}
	return "", errors.New("todo not found")
}

func (s *StubTodoStore) AddTodo(description string) {
	s.todos[1] = models.NewTodo(1, description, false)
}
