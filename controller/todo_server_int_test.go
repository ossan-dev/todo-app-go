package controller_test

import (
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
	}

	res := httptest.NewRecorder()
	server.ServeHTTP(res, util.NewPostTodoReq(t, "/api/todos", &wantedTodos[0]))

	assert.Equal(t, http.StatusCreated, res.Code)
}
