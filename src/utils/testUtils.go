package utils

import (
	"reflect"
	"testing"

	"todo-app-go.com/v1/src/models"
)

func AssertTodosEqual(t *testing.T, got, want models.Todo) {
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

func AssertCollectionsEqual(t *testing.T, got, want []models.Todo) {
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
