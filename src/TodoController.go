package src

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-elastic/src/elastic"
	"go-elastic/src/response"
	"log"
	"net/http"
)

const TodosIndex = "todos"

func EnrichRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/todo").Subrouter()
	subRouter.HandleFunc("", createTodo).Methods(http.MethodPost)
	subRouter.HandleFunc("/{id}", removeTodo).Methods(http.MethodDelete)
	subRouter.HandleFunc("/{id}", putTodo).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", readTodo).Methods(http.MethodGet)

	todosRouter := router.PathPrefix("/todos").Subrouter()
	todosRouter.HandleFunc("", searchTodos).Methods(http.MethodGet)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	responseId, err := elastic.GetElasticInstance().InsertDocument("", &todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	responseData := response.NewRequestResponse(http.StatusCreated, map[string]string{"id": responseId, "status": "todo created"})
	response.WriteResponse(responseData, w)

}

func readTodo(w http.ResponseWriter, r *http.Request) {
	todoID, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Problem retrieving [id] from URL")
		return
	}

	var todo Todo
	if err := elastic.GetElasticInstance().SearchDocument(todoID, &todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := response.NewRequestResponse(http.StatusOK, todo)
	response.WriteResponse(responseData, w)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Problem retrieving [id] from URL")
		return
	}

	elastic.GetElasticInstance().DeleteDocument(todoID, TodosIndex)
	responseData := response.NewRequestResponse(http.StatusOK, nil)
	response.WriteResponse(responseData, w)
}

func putTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	todoID, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Problem retrieving [id] from URL")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseId, err := elastic.GetElasticInstance().InsertDocument(todoID, &todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	responseData := response.NewRequestResponse(
		http.StatusOK,
		map[string]string{
			"id":     responseId,
			"status": "todo updated",
		},
	)

	response.WriteResponse(responseData, w)
}

func searchTodos(w http.ResponseWriter, r *http.Request) {
	log.Printf("Search todo log")
}
