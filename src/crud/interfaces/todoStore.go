package interfaces

type TodoStore interface {
	GetTodoById(id int) (string, error)
	AddTodo(description string)
}
