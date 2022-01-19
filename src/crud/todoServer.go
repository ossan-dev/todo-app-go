package crud

import (
	"fmt"
	"net/http"
	"strings"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	if id == "1" {
		fmt.Fprintf(w, `{"id": 1, "description": "FirstTodo", "isCompleted": false}`)
		return
	}

	if id == "2" {
		fmt.Fprintf(w, `{"id": 2, "description": "SecondTodo", "isCompleted": true}`)
		return
	}
}
