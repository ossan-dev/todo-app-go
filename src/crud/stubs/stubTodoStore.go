package stubs

import "todo-app-go.com/v1/src/models"

type StubTodoStore struct {
	Todos map[int]models.Todo
}

func (s *StubTodoStore) GetTodoById(id int) string {
	todo := s.Todos[id]
	return todo.Description
}

func (s *StubTodoStore) AddTodo(description string) {
	s.Todos[1] = models.NewTodo(1, description, false)
}
