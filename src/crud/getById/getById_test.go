package getbyid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/model"
	"todo-app-go.com/v1/src/crud/servers"
	"todo-app-go.com/v1/util"
)

func TestGetByIdEndpoint(t *testing.T) {
	store := database.NewStubTodoStore(
		&map[int]model.Todo{
			1: model.NewTodo(1, "FirstTodo", false),
			2: model.NewTodo(2, "SecondTodo", false),
		},
	)
	server := servers.NewTodoServer(store)

	t.Run("first todo", func(t *testing.T) {
		request := newGetToDoByIdReq(1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		util.AssertStatusCode(t, response.Code, http.StatusOK)
		util.AssertResponseBody(t, response.Body.String(), "FirstTodo")
	})

	t.Run("second todo", func(t *testing.T) {
		request := newGetToDoByIdReq(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		util.AssertStatusCode(t, response.Code, http.StatusOK)
		util.AssertResponseBody(t, response.Body.String(), "SecondTodo")
	})

	t.Run("returns 404 on missing todos", func(t *testing.T) {
		request := newGetToDoByIdReq(4)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		util.AssertResponseBody(t, response.Body.String(), "")
		util.AssertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func newGetToDoByIdReq(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/todos/%d", id), nil)
	return req
}
