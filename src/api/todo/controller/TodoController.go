package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/strikersk/go-elastic/src/api/todo/entity"
	"github.com/strikersk/go-elastic/src/api/todo/repository"
	"github.com/strikersk/go-elastic/src/response"
	"log"
	"net/http"
)

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
	var todo entity.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		res := response.NewRequestResponse(http.StatusInternalServerError, err)
		response.WriteResponse(res, w)
		return
	}

	responseId, err := repository.TodoRepository.InsertDocument("", todo)
	if err != nil {
		res := response.NewRequestResponse(http.StatusInternalServerError, err)
		response.WriteResponse(res, w)
		return
	}

	res := response.NewRequestResponse(http.StatusCreated, map[string]string{"id": responseId, "status": "todo created"})
	response.WriteResponse(res, w)
}

func readTodo(w http.ResponseWriter, r *http.Request) {
	todoID, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Problem retrieving [id] from URL")
		return
	}

	todo, err := repository.TodoRepository.SearchDocument(todoID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := response.NewRequestResponse(http.StatusOK, todo)
	response.WriteResponse(res, w)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Problem retrieving [id] from URL")
		return
	}

	_ = repository.TodoRepository.DeleteDocument(todoID)
	res := response.NewRequestResponse(http.StatusOK, nil)
	response.WriteResponse(res, w)
}

func putTodo(w http.ResponseWriter, r *http.Request) {
	var todo entity.Todo
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
	responseId, err := repository.TodoRepository.InsertDocument(todoID, todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	res := response.NewRequestResponse(
		http.StatusOK,
		map[string]string{
			"id":     responseId,
			"status": "todo updated",
		},
	)

	response.WriteResponse(res, w)
}

func searchTodos(w http.ResponseWriter, r *http.Request) {
	log.Printf("Search todo log")
}
