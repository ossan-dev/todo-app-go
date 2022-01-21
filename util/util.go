package util

import (
	"sort"
	"testing"

	"todo-app-go.com/v1/model"
)

func SortTodoSliceById(todos []model.Todo) {
	sort.Slice(todos, func(t1, t2 int) bool {
		return todos[t1].Id < todos[t2].Id
	})
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
