package util

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sort"
	"testing"

	"todo-app-go.com/v1/model"
)

var JsonContentType string = "application/json"

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

func NewPostReq(t *testing.T, url string, todo *model.Todo) *http.Request {
	t.Helper()
	buf, err := json.Marshal(todo)
	if err != nil {
		t.Fatalf("could not serialize model %v, err %v", todo, err)
	}

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(buf))
	return req
}

func NewGetTodoReq(t *testing.T, url string) *http.Request {
	t.Helper()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return req
}
