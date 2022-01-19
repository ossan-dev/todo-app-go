package main

import (
	"log"
	"net/http"

	"todo-app-go.com/v1/src/crud"
)

func main() {
	handler := http.HandlerFunc(crud.TodoServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
