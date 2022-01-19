package crud

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type TodoServer struct {
	Store TodoStore
}

func (t *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/todos/"))
	fmt.Fprintf(w, t.Store.GetTodoById(id))
}

type TodoStore interface {
	GetTodoById(id int) string
}

func GetTodoById(id int) string {
	if id == 1 {
		return `{"id": 1, "description": "FirstTodo", "isCompleted": false}`
	}

	if id == 2 {
		return `{"id": 2, "description": "SecondTodo", "isCompleted": true}`
	}

	return ""
}
