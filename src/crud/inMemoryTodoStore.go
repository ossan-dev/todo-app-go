package crud

type InMemoryTodoStore struct{}

func (i *InMemoryTodoStore) GetTodoById(id int) string {
	return "FirstTodo"
}

func (i *InMemoryTodoStore) AddTodo(description string) {}
