package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-app-go.com/v1/controller"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/model"
)

func TestGetAll(t *testing.T) {
	todo := model.NewTodo(1, "FirstTodo", false)
	store := database.NewStubTodoStore(
		&map[int]model.Todo{
			1: todo,
		},
	)
	server := controller.NewTodoServer(&store)

	t.Run("it returns 200 on /api/todos", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		var got []model.Todo

		err := json.NewDecoder(res.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Todo, %v", res.Body, err)
		}

		assert.Equal(t, res.Code, http.StatusOK)
	})
}
