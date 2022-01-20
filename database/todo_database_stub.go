package database

import (
	"errors"

	"todo-app-go.com/v1/error_handler"
	"todo-app-go.com/v1/model"
)

type StubTodoStore struct {
	todos map[int]model.Todo
}

func NewStubTodoStore(todos *map[int]model.Todo) *StubTodoStore {
	return &StubTodoStore{todos: *todos}
}

func (s *StubTodoStore) GetTodoById(id int) (*model.Todo, error) {
	if val, ok := s.todos[id]; ok {
		todo := val
		return &todo, nil
	}
	return nil, error_handler.ErrNotFound
}

func (s *StubTodoStore) GetAllTodos() []model.Todo {
	result := make([]model.Todo, 0)
	for _, val := range s.todos {
		result = append(result, val)
	}
	return result
}

func (s *StubTodoStore) AddTodo(description string) (int, error) {
	if description == "" {
		return 0, errors.New("todo instance not correct")
	}
	s.todos[1] = model.NewTodo(1, description, false)
	return 1, nil
}
