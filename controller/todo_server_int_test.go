package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-app-go.com/v1/controller"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/model"
	"todo-app-go.com/v1/util"
)

func TestSavingTodosAndRetrievingThem(t *testing.T) {
	store := database.NewInMemoryTodoStore()
	server := controller.NewTodoServer(&store)
	wantedTodos := []model.Todo{
		model.NewTodo(1, "First Todo", false),
		model.NewTodo(2, "Second Todo", true),
	}

	t.Run("post todos", func(t *testing.T) {
		res := httptest.NewRecorder()

		server.ServeHTTP(res, util.NewPostReq(t, "/api/todos", &wantedTodos[0]))
		assert.Equal(t, http.StatusCreated, res.Code)

		server.ServeHTTP(res, util.NewPostReq(t, "/api/todos", &wantedTodos[1]))
		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("get todos", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, util.NewGetTodoReq(t, "/api/todos"))

		var todos []model.Todo
		json.NewDecoder(res.Body).Decode(&todos)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, wantedTodos, todos)
	})
}

func TestSavingTodoAndRetrievingIt(t *testing.T) {
	store := database.NewInMemoryTodoStore()
	server := controller.NewTodoServer(&store)
	wantedTodo := model.NewTodo(1, "First Todo", false)

	t.Run("post todo", func(t *testing.T) {
		res := httptest.NewRecorder()

		server.ServeHTTP(res, util.NewPostReq(t, "/api/todos", &wantedTodo))
		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("get todos", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, util.NewGetTodoReq(t, fmt.Sprintf("/api/todos/%d", wantedTodo.Id)))

		var todos model.Todo
		json.NewDecoder(res.Body).Decode(&todos)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, wantedTodo, todos)
	})
}
