package src

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}
	ESConfiguration.createData(todo)
}

func readTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("property [id] does no exist")
		return
	}

	persistedTodo, err := ESConfiguration.getTodo(todoID).MarshalItem()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(persistedTodo)
}
