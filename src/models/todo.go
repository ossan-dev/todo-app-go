package models

// TODO: make these props not visible to others
type Todo struct {
	Id          int
	Description string
	IsCompleted bool
}

func NewTodo(id int, description string, isCompleted bool) Todo {
	// TODO: make isCompleted optional
	return Todo{id, description, isCompleted}
}
