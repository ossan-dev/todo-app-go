package crud

import "todo-app-go.com/v1/src/models"

type StubTodoStore struct {
	todos map[int]models.Todo
}

func (s *StubTodoStore) GetTodoById(id int) string {
	todo := s.todos[id]
	return todo.Description
}
