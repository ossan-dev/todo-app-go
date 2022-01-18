package test

import (
	"testing"

	"todo-app-go.com/v1/src/models"
)

func TestGetAll(t *testing.T) {
	todos := []models.Todo{
		{1, "FirstTodo", false},
		{2, "SecondTodo", false},
		{3, "ThirdTodo", false},
	}

	todoManager := &models.TodoManager{todos}

	got := todoManager.GetAllTodos()

	if len(got) != len(todos) {
		t.Errorf("got %d TODOS but wanted %d TODOS", len(got), len(todos))
	}
}
