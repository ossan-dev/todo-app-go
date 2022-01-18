package crud

import "todo-app-go.com/v1/src/models"

type TodoManager struct {
	Todos []models.Todo
}

func (t *TodoManager) GetAllTodos() []models.Todo {
	return t.Todos
}

func (t *TodoManager) GetById(id int) models.Todo {
	return t.Todos[id-1]
}
