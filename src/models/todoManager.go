package models

type TodoManager struct {
	Todos []Todo
}

func (t *TodoManager) GetAllTodos() []Todo {
	return t.Todos
}

func (t *TodoManager) GetById(id int) Todo {
	return Todo{}
}
