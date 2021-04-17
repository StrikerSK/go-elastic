package src

import (
	"github.com/gorilla/mux"
)

const todosIndex = "custom_todos"

func EnrichRouter(mainRouter *mux.Router) {

	exampleRouter := mainRouter.PathPrefix("/example").Subrouter()
	exampleRouter.HandleFunc("/index", createExampleIndexBody).Methods("GET")
	exampleRouter.HandleFunc("/type", createExampleType).Methods("GET")

	todoRouter := mainRouter.PathPrefix("/todo").Subrouter()

}
