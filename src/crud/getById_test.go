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
		got, err := todoManager.GetById(1)

		assertTodosEqual(t, got, &todos[0])
		assertNoError(t, err)
	})

	t.Run("todo not present in collection", func(t *testing.T) {
		got, err := todoManager.GetById(4)

		assertTodosEqual(t, got, nil)
		assertError(t, err, ErrTodoNotFound)
	})
}

func assertTodosEqual(t *testing.T, got, want *models.Todo) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error but didn't find one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got an error but didn't want one")
	}
}
