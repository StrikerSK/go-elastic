package exampleType

import (
	"github.com/gorilla/mux"
	"github.com/strikersk/go-elastic/src/api/todo"
	"github.com/strikersk/go-elastic/src/elastic/body"
	"github.com/strikersk/go-elastic/src/response"
	"net/http"
)

func EnrichRouterWithExamples(mainRouter *mux.Router) {
	exampleRouter := mainRouter.PathPrefix("/example").Subrouter()
	exampleRouter.HandleFunc("/index", createExampleIndexBody).Methods(http.MethodGet)
	exampleRouter.HandleFunc("/type", createExampleType).Methods(http.MethodGet)
}

func createExampleType(w http.ResponseWriter, r *http.Request) {
	customTodo := todo.Todo{
		Name:        "Example Create Todo",
		Description: "Example Create Todo",
		Done:        false,
	}

	res := response.NewRequestResponse(http.StatusOK, customTodo)
	response.WriteResponse(res, w)
}

func createExampleIndexBody(w http.ResponseWriter, r *http.Request) {
	res := response.NewRequestResponse(http.StatusOK, body.NewDefaultElasticBody(*body.CreateMappingMap(exampleStruct{})))
	response.WriteResponse(res, w)
}
