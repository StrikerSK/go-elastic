package src

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const TODOS_INDEX = "custom_todos"
const HOST_URL = "http://localhost:9200"

func EnrichRouter(mainRouter *mux.Router) {

	todoRouter := mainRouter.PathPrefix("/todo").Subrouter()
	todoRouter.HandleFunc("", createTodo).Methods("POST")
	todoRouter.HandleFunc("/{id}", removeTodo).Methods("DELETE")
	todoRouter.HandleFunc("/{id}", readTodo).Methods("GET")

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
	responseData := createData(todo)
	outputData, _ := json.Marshal(responseData)
	w.WriteHeader(responseData.Code)
	_, _ = w.Write(outputData)

}

func readTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Problem retrieving [id] from URL")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseData := getTodo(todoID)
	outputData, _ := json.Marshal(responseData)
	w.WriteHeader(responseData.Code)
	_, _ = w.Write(outputData)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Problem retrieving [id] from URL")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseData := deleteTodo(todoID)
	outputData, _ := json.Marshal(responseData)
	w.WriteHeader(responseData.Code)
	_, _ = w.Write(outputData)
}
