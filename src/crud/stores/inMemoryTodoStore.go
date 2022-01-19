package stores

import "todo-app-go.com/v1/src/models"

type InMemoryTodoStore struct {
	todos map[int]models.Todo
}

func NewInMemoryTodoStore() *InMemoryTodoStore {
	todos := make(map[int]models.Todo, 0)
	return &InMemoryTodoStore{
		todos: todos,
	}
}

func (i *InMemoryTodoStore) GetTodoById(id int) (string, error) {
	return i.todos[id].Description, nil
}

func (i *InMemoryTodoStore) AddTodo(description string) {
	maxKey := 0
	for key := range i.todos {
		if key > maxKey {
			maxKey = key
		}
	}

	i.todos[maxKey+1] = models.NewTodo(maxKey+1, description, false)
}
