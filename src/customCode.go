package src

import "net/http"

func createExampleType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customTodo := Todo{
		Name:        "Create todo",
		Description: "Describe todo",
		Done:        false,
	}

	if _, err := w.Write(customTodo.marshalTodo()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func createExampleIndexBody(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(CreateTodoIndexBody())
}
