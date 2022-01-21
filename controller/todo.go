package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/error_handler"
	"todo-app-go.com/v1/model"
)

type TodoController struct {
	store database.TodoStore
}

// GET /api/todos
// get all todos
func (tc *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	todos := tc.store.GetAllTodos()
	json.NewEncoder(w).Encode(todos)
	w.WriteHeader(http.StatusOK)
}

// GET /api/todos/:id
// get todo by id
func (tc *TodoController) GetTodoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, _ := strconv.Atoi(idParam)
	w.Header().Set("content-type", "application/json")
	todo, _ := tc.store.GetTodoById(id)
	json.NewEncoder(w).Encode(todo)
	w.WriteHeader(http.StatusOK)
}

// POST /api/todos
// create a new todo
func (tc *TodoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	parseErr := json.NewDecoder(r.Body).Decode(&todo)
	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := error_handler.TodoAppError{Msg: error_handler.INVALID_BODY_REQUEST_MSG, Code: error_handler.BAD_REQUEST_CODE, Details: []string{parseErr.Error()}}
		json.NewEncoder(w).Encode(err)
		return
	}
	tc.store.AddTodo(todo)
	w.WriteHeader(http.StatusCreated)
}
