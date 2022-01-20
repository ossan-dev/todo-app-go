package database

import "todo-app-go.com/v1/model"

type InMemoryTodoStore struct {
	todos map[int]model.Todo
}

type TodoStore interface {
	GetTodoById(id int) (string, error)
	AddTodo(description string) int
}

func NewInMemoryTodoStore() *InMemoryTodoStore {
	todos := make(map[int]model.Todo, 0)
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

	i.todos[maxKey+1] = model.NewTodo(maxKey+1, description, false)
}
