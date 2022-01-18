package test

import "testing"

func TestGetAll(t *testing.T) {
	todos := []Todo{
		{1, "FirstTodo", false},
		{2, "SecondTodo", false},
		{3, "ThirdTodo", false},
	}

	todoManager := &TodoManager{todos}

	got := todoManager.GetAllTodos()

	if len(got) != len(todos) {
		t.Errorf("got %d TODOS but wanted %d TODOS", len(got), len(todos))
	}
}
