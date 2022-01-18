package models

type TodoManager struct {
	Todos []Todo
}

func (t *TodoManager) GetAllTodos() []Todo {
	return t.Todos
}
