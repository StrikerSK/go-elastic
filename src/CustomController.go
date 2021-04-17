package src

import (
	"github.com/gorilla/mux"
)

const TODOS_INDEX = "custom_todos"
const HOST_URL = "http://localhost:9200"

func EnrichRouter(mainRouter *mux.Router) {

	exampleRouter := mainRouter.PathPrefix("/example").Subrouter()
	exampleRouter.HandleFunc("/index", createExampleIndexBody).Methods("GET")
	exampleRouter.HandleFunc("/type", createExampleType).Methods("GET")

	todoRouter := mainRouter.PathPrefix("/todo").Subrouter()
	todoRouter.HandleFunc("", createTodo).Methods("POST")

}
