package database

import (
	"todo-app-go.com/v1/error_handler"
	"todo-app-go.com/v1/model"
)

type InMemoryTodoStore struct {
	todos map[int]model.Todo
}

type TodoStore interface {
	GetTodoById(id int) (*model.Todo, error)
	GetAllTodos() []model.Todo
	GetByStatus(status bool) []model.Todo
	AddTodo(todo model.Todo) (int, error)
	UpdateTodo(todo model.Todo) (*int, error)
	DeleteById(id int) (int, error)
}

func NewInMemoryTodoStore() InMemoryTodoStore {
	todos := make(map[int]model.Todo, 0)
	return InMemoryTodoStore{
		todos: todos,
	}
}

func (i *InMemoryTodoStore) GetTodoById(id int) (*model.Todo, error) {
	var todo model.Todo
	return &todo, nil
}

func (i *InMemoryTodoStore) GetAllTodos() []model.Todo {
	todos := make([]model.Todo, 0)
	for _, val := range i.todos {
		todos = append(todos, val)
	}
	return todos
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

func (i *InMemoryTodoStore) UpdateTodo(todo model.Todo) (*int, error) {
	for key, val := range i.todos {
		if val.Id == todo.Id {
			i.todos[key] = val
			return &todo.Id, nil
		}
	}
	return nil, error_handler.ErrNotFound
}

func (i *InMemoryTodoStore) DeleteById(id int) (int, error) {
	if _, ok := i.todos[id]; ok {
		delete(i.todos, id)
		return 1, nil
	}
	return 0, error_handler.ErrNotFound
}
