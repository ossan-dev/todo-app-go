package servers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"todo-app-go.com/v1/src/crud/interfaces"
)

type TodoServer struct {
	store interfaces.TodoStore
}

func NewTodoServer(todoStore interfaces.TodoStore) *TodoServer {
	return &TodoServer{todoStore}
}

func (t *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/todos/"))
	switch r.Method {
	case http.MethodPost:
		t.todoCreation(w, r)
	case http.MethodGet:
		t.showDescription(w, id)
	}
}

func (t *TodoServer) showDescription(w http.ResponseWriter, todoId int) {
	todo := t.store.GetTodoById(todoId)

	if todo == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, todo)
}

func (t *TodoServer) todoCreation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	reqBody, _ := ioutil.ReadAll(r.Body)
	t.store.AddTodo(string(reqBody))
}
