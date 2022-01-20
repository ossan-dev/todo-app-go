package util

import (
	"reflect"
	"sort"
	"testing"

	"todo-app-go.com/v1/model"
)

func SortTodoSliceById(todos []model.Todo) {
	sort.Slice(todos, func(t1, t2 int) bool {
		return todos[t1].Id < todos[t2].Id
	})
}

func AssertTodosEqual(t *testing.T, got, want model.Todo) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func AssertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error but didn't find one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q but want %q", got, want)
	}
}

func AssertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got an error but didn't want one")
	}
}

func AssertCollectionsEqual(t *testing.T, got, want []model.Todo) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func AssertStatusCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
