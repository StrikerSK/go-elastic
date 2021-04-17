package src

import (
	"log"
	"net/http"
)

func createExampleType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customTodo := Todo{
		Name:        "Create todo",
		Description: "Describe todo",
		Done:        false,
	}

	marshalled, err := customTodo.MarshalItem()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	if _, err = w.Write(marshalled); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}

func createExampleIndexBody(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(CreateTodoIndexBody())
}
