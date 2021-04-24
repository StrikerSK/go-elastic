package src

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-elastic/src/response"
	"log"
	"net/http"
)

const TodosIndex = "custom_todos"
const HostUrl = "http://localhost:9200"

func EnrichRouter(mainRouter *mux.Router) {

	todoRouter := mainRouter.PathPrefix("/todo").Subrouter()
	todoRouter.HandleFunc("", createTodo).Methods("POST")
	todoRouter.HandleFunc("/{id}", removeTodo).Methods("DELETE")
	todoRouter.HandleFunc("/{id}", putTodo).Methods("PUT")
	todoRouter.HandleFunc("/{id}", readTodo).Methods("GET")

	todosRouter := mainRouter.PathPrefix("/todos").Subrouter()
	todosRouter.HandleFunc("", searchTodos).Methods("GET")

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

	responseId, err := ESConfiguration.insertToIndex("", &todo, TodosIndex)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	responseData := response.RequestResponse{
		Data:   responseId,
		Status: "Todo created",
		Code:   http.StatusCreated,
	}
	outputData, _ := json.Marshal(responseData)

	w.WriteHeader(responseData.Code)
	_, _ = w.Write(outputData)

}

func readTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Problem retrieving [id] from URL")
		return
	}

	var todo Todo
	if err := ESConfiguration.searchTodos(TodosIndex, todoID, &todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseData := response.RequestResponse{
		Data:   todo,
		Status: "Todo found",
		Code:   http.StatusOK,
	}
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

	ESConfiguration.deleteItem(TodosIndex, todoID)

	w.Header().Set("Content-Type", "application/json")
	responseData := response.RequestResponse{
		Data:   nil,
		Status: "Todo deleted",
		Code:   200,
	}

	outputData, _ := json.Marshal(responseData)
	w.WriteHeader(responseData.Code)
	_, _ = w.Write(outputData)
}

func putTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	vars := mux.Vars(r)
	todoID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Problem retrieving [id] from URL")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseId, err := ESConfiguration.insertToIndex(todoID, &todo, TodosIndex)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	responseData := response.RequestResponse{
		Data:   responseId,
		Status: "Todo updated",
		Code:   http.StatusOK,
	}
	outputData, _ := json.Marshal(responseData)
	w.WriteHeader(responseData.Code)
	_, _ = w.Write(outputData)
}

func searchTodos(w http.ResponseWriter, r *http.Request) {
	log.Printf("Search todo log")
}
