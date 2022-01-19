package crud

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestGetById(t *testing.T) {
	store := StubTodoStore{
		map[int]models.Todo{
			1: models.NewTodo(1, "FirstTodo", false),
			2: models.NewTodo(2, "SecondTodo", false),
		},
	}

	t.Run("todo present in collection", func(t *testing.T) {
		got := store.GetTodoById(1)

		utils.AssertResponseBody(t, got, store.todos[1].Description)
		// utils.AssertNoError(t, err)
	})

	// t.Run("todo not present in collection", func(t *testing.T) {
	// 	got, err := store.GetTodoById(4)

	// 	utils.AssertTodosEqual(t, got, nil)
	// 	utils.AssertError(t, err, ErrTodoNotFound)
	// })
}

func TestGetByIdEndpoint(t *testing.T) {
	store := StubTodoStore{
		map[int]models.Todo{
			1: models.NewTodo(1, "FirstTodo", false),
			2: models.NewTodo(2, "SecondTodo", false),
		},
	}
	server := &TodoServer{&store}

	t.Run("first todo", func(t *testing.T) {
		request := newGetToDoByIdReq(1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), store.todos[1].Description)
	})

	t.Run("second todo", func(t *testing.T) {
		request := newGetToDoByIdReq(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), store.todos[2].Description)
	})
}

func newGetToDoByIdReq(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/todos/%d", id), nil)
	return req
}
