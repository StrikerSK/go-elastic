package main

import (
	"github.com/gorilla/mux"
	"go-elastic/src/api/exampleType"
	"go-elastic/src/api/todo"
	"go-elastic/src/elastic"
	"log"
	"net/http"
	"os"
)

func init() {
	elastic.GetElasticInstance().InitializeIndex(todo.TodosIndex, todo.CreateTodoIndexBody())
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	todo.EnrichRouter(router)
	exampleType.EnrichRouterWithExamples(router)

	log.Println("Listening")
	log.Println(http.ListenAndServe(":"+resolvePort(), router))
}

func createServer() *http.Server {
	router := mux.NewRouter().StrictSlash(true)
	todo.EnrichRouter(router)
	exampleType.EnrichRouterWithExamples(router)

	return &http.Server{
		Addr:    resolvePort(),
		Handler: router,
	}
}

func resolvePort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		log.Printf("Cannot retrieve port number using default value\n")
		port = "5000"
	}
	return
}
