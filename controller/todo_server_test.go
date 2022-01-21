package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-app-go.com/v1/controller"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/model"
)

func TestGetAll(t *testing.T) {
	wantedTodos := []model.Todo{
		model.NewTodo(1, "FirstTodo", false),
		model.NewTodo(2, "SecondTodo", false),
	}
	store := database.NewStubTodoStore(
		&map[int]model.Todo{
			1: wantedTodos[0],
			2: wantedTodos[1],
		},
	)
	server := controller.NewTodoServer(&store)

	t.Run("it todos as JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		var got []model.Todo

		err := json.NewDecoder(res.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Todo, %v", res.Body, err)
		}

		assert.Equal(t, res.Code, http.StatusOK)

		if !reflect.DeepEqual(got, wantedTodos) {
			t.Errorf("got %v but want %v", got, wantedTodos)
		}
	})
}
