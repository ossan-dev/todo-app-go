package crud

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

// func TestAdd(t *testing.T) {
// 	store := StubTodoStore{
// 		map[int]models.Todo{
// 			1: models.NewTodo(1, "FirstTodo", false),
// 			2: models.NewTodo(2, "SecondTodo", false),
// 		},
// 	}

// 	todoManager := &TodoManager{todos}

// 	todo := models.NewTodo(4, "Fourth todo", false)
// 	todoManager.Add(todo)

// 	got := .GetById(todo.Id)

// 	utils.AssertNoError(t, err)
// 	utils.AssertTodosEqual(t, got, &todo)
// }

func TestAddTodo(t *testing.T) {
	store := StubTodoStore{
		map[int]models.Todo{},
	}

	server := &TodoServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request := newPostTodoReq("Example Todo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertStatusCode(t, response.Code, http.StatusAccepted)

		if len(store.todos) != 1 {
			t.Errorf("got %d calls to AddTodo but want %d", len(store.todos), 1)
		}
	})
}

func newPostTodoReq(description string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/api/todos", strings.NewReader(description))
	return req
}
