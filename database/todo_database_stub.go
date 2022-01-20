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

func (s *StubTodoStore) GetByStatus(status bool) []model.Todo {
	result := make([]model.Todo, 0)
	for _, val := range s.todos {
		if val.IsCompleted == status {
			result = append(result, val)
		}
	}
	return result
}

func (s *StubTodoStore) AddTodo(todo model.Todo) (int, error) {
	if todo.Description == "" {
		return 0, error_handler.ErrTodoNotValid
	}
	s.todos[1] = todo
	return 1, nil
}
