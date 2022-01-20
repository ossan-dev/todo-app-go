package database

import (
	"todo-app-go.com/v1/error_handler"
	"todo-app-go.com/v1/model"
)

type StubTodoStore struct {
	todos map[int]model.Todo
}

func NewStubTodoStore(todos *map[int]model.Todo) *StubTodoStore {
	return &StubTodoStore{todos: *todos}
}

func (s *StubTodoStore) GetTodoById(id int) (string, error) {
	if val, ok := s.todos[id]; ok {
		todo := val
		return todo.Description, nil
	}
	return "", error_handler.ErrNotFound
}

func (s *StubTodoStore) AddTodo(description string) int {
	s.todos[1] = model.NewTodo(1, description, false)
	return 1
}
