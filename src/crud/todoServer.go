package crud

import (
	"fmt"
	"net/http"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"id": 1, "description": "FirstTodo", "isCompleted": false}`)
}
