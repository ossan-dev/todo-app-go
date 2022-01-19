package crud

import "todo-app-go.com/v1/src/models"

type InMemoryTodoStore struct {
	store map[int]models.Todo
}

func NewInMemoryTodoStore() *InMemoryTodoStore {
	store := make(map[int]models.Todo, 0)
	return &InMemoryTodoStore{
		store: store,
	}
}

func (i *InMemoryTodoStore) GetTodoById(id int) string {
	return i.store[id].Description
}

func (i *InMemoryTodoStore) AddTodo(description string) {
	maxKey := 0
	for key := range i.store {
		if key > maxKey {
			maxKey = key
		}
	}

	i.store[maxKey+1] = models.NewTodo(maxKey+1, description, false)
}
