package crud

import "todo-app-go.com/v1/src/models"

type TodoManager struct {
	Todos []models.Todo
}

func (t *TodoManager) GetAllTodos() []models.Todo {
	return t.Todos
}

func (t *TodoManager) GetById(id int) models.Todo {
	for _, todo := range t.Todos {
		if todo.Id == id {
			return todo
		}
	}
	return models.Todo{}
}
