package crud

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type TodoServer struct {
	Store TodoStore
}

func (t *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		t.todoCreation(w, r)
	case http.MethodGet:
		t.showDescription(w, r)
	}
}

func (t *TodoServer) showDescription(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/todos/"))

	todo := t.Store.GetTodoById(id)

	if todo == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, todo)
}

func (t *TodoServer) todoCreation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	reqBody, _ := ioutil.ReadAll(r.Body)
	t.Store.AddTodo(string(reqBody))
}
