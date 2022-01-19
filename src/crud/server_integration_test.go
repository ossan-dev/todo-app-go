package crud

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-app-go.com/v1/src/utils"
)

func TestSavingTodosAndRetrievingThem(t *testing.T) {
	store := NewInMemoryTodoStore()
	server := TodoServer{store}
	todo := "Example"

	server.ServeHTTP(httptest.NewRecorder(), newPostTodoReq(todo))
	server.ServeHTTP(httptest.NewRecorder(), newPostTodoReq(todo))
	server.ServeHTTP(httptest.NewRecorder(), newPostTodoReq(todo))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetToDoByIdReq(1))
	utils.AssertStatusCode(t, response.Code, http.StatusOK)
	utils.AssertResponseBody(t, response.Body.String(), "Example")

	response = httptest.NewRecorder()
	server.ServeHTTP(response, newGetToDoByIdReq(2))
	utils.AssertStatusCode(t, response.Code, http.StatusOK)
	utils.AssertResponseBody(t, response.Body.String(), "Example")

	response = httptest.NewRecorder()
	server.ServeHTTP(response, newGetToDoByIdReq(3))
	utils.AssertStatusCode(t, response.Code, http.StatusOK)
	utils.AssertResponseBody(t, response.Body.String(), "Example")
}
