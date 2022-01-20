package database

import "todo-app-go.com/v1/model"

type InMemoryTodoStore struct {
	todos map[int]model.Todo
}

type TodoStore interface {
	GetTodoById(id int) (model.Todo, error)
	GetAllTodos() []model.Todo
	GetByStatus(status bool) []model.Todo
	AddTodo(todo model.Todo) (int, error)
}

func NewInMemoryTodoStore() *InMemoryTodoStore {
	todos := make(map[int]model.Todo, 0)
	return &InMemoryTodoStore{
		todos: todos,
	}
}

func (i *InMemoryTodoStore) GetTodoById(id int) (model.Todo, error) {
	return i.todos[id], nil
}

func (i *InMemoryTodoStore) GetAllTodos() []model.Todo {
	return []model.Todo{}
}

func (i *InMemoryTodoStore) GetByStatus(status bool) []model.Todo {
	return []model.Todo{}
}

func (i *InMemoryTodoStore) AddTodo(todo model.Todo) (int, error) {
	maxKey := 0
	for key := range i.todos {
		if key > maxKey {
			maxKey = key
		}
	}

	i.todos[maxKey+1] = todo
	return 1, nil
}
