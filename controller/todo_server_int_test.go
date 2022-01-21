package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-app-go.com/v1/controller"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/model"
)

func TestSavingTodosAndRetrievingThem(t *testing.T) {
	store := database.NewInMemoryTodoStore()
	server := controller.NewTodoServer(&store)
	wantedTodos := []model.Todo{
		model.NewTodo(1, "First Todo", false),
	}

	res := httptest.NewRecorder()
	server.ServeHTTP(res, newPostTodoReq(t, &wantedTodos[0]))

	assert.Equal(t, http.StatusCreated, res.Code)
}

func newPostTodoReq(t *testing.T, todo *model.Todo) *http.Request {
	t.Helper()
	buf, err := json.Marshal(todo)
	if err != nil {
		t.Fatalf("could not serialize model %v, err %v", todo, err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/api/todos", bytes.NewBuffer(buf))
	return req
}
