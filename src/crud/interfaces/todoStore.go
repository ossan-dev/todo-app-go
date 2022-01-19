package interfaces

type TodoStore interface {
	GetTodoById(id int) string
	AddTodo(description string)
}
