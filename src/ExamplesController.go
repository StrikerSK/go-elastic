package src

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func EnrichRouterWithExamples(mainRouter *mux.Router) {

	exampleRouter := mainRouter.PathPrefix("/example").Subrouter()
	exampleRouter.HandleFunc("/index", createExampleIndexBody).Methods("GET")
	exampleRouter.HandleFunc("/type", createExampleType).Methods("GET")

}

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
