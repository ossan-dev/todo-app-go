package getbyid

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-app-go.com/v1/src/crud/servers"
	"todo-app-go.com/v1/src/crud/stubs"
	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestGetById(t *testing.T) {
	store := stubs.NewStubTodoStore(
		&map[int]models.Todo{
			1: models.NewTodo(1, "FirstTodo", false),
			2: models.NewTodo(2, "SecondTodo", false),
		},
	)

	t.Run("todo present in collection", func(t *testing.T) {
		got, err := store.GetTodoById(1)

		utils.AssertResponseBody(t, got, "FirstTodo")
		utils.AssertNoError(t, err)
	})

	t.Run("todo not present in collection", func(t *testing.T) {
		got, err := store.GetTodoById(4)

		utils.AssertResponseBody(t, got, "")
		utils.AssertError(t, err, errors.New("todo not found"))
	})
}

func TestGetByIdEndpoint(t *testing.T) {
	store := stubs.NewStubTodoStore(
		&map[int]models.Todo{
			1: models.NewTodo(1, "FirstTodo", false),
			2: models.NewTodo(2, "SecondTodo", false),
		},
	)
	server := servers.NewTodoServer(store)

	t.Run("first todo", func(t *testing.T) {
		request := newGetToDoByIdReq(1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatusCode(t, response.Code, http.StatusOK)
		utils.AssertResponseBody(t, response.Body.String(), "FirstTodo")
	})

	t.Run("second todo", func(t *testing.T) {
		request := newGetToDoByIdReq(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatusCode(t, response.Code, http.StatusOK)
		utils.AssertResponseBody(t, response.Body.String(), "SecondTodo")
	})

	t.Run("returns 404 on missing todos", func(t *testing.T) {
		request := newGetToDoByIdReq(4)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func newGetToDoByIdReq(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/todos/%d", id), nil)
	return req
}
