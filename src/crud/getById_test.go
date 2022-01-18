package crud

import (
	"reflect"
	"testing"

	"todo-app-go.com/v1/src/models"
)

func TestGetById(t *testing.T) {
	todos := []models.Todo{
		{1, "FirstTodo", false},
		{2, "SecondTodo", false},
		{3, "ThirdTodo", false},
	}

	todoManager := &TodoManager{todos}

	t.Run("todo present in collection", func(t *testing.T) {
		assertTodosEqual(t, todoManager.GetById(1), todos[0])
	})

	t.Run("todo not present in collection", func(t *testing.T) {
		_, err := todoManager.GetById(4)
		if err != nil {
			t.Errorf("expected an error but didn't find one")
		}
	})
}

func assertTodosEqual(t *testing.T, got, want models.Todo) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
