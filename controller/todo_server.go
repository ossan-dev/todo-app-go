package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/model"
)

type TodoServer struct {
	Store database.TodoStore
}

func NewTodoServer(todoStore *database.TodoStore) *TodoServer {
	return &TodoServer{*todoStore}
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
	todo, err := t.Store.GetTodoById(todoId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, "%v", todo)
}

func (t *TodoServer) todoCreation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo model.Todo
	json.Unmarshal(reqBody, &todo)
	t.Store.AddTodo(todo)
}
