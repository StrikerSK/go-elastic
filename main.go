package main

import (
	"github.com/gorilla/mux"
	"go-elastic/src"
	"log"
	"net/http"
	"os"
)

func init() {
	src.GetElasticInstance().CreateIndex(src.TodosIndex, src.CreateTodoIndexBody())
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	src.EnrichRouter(router)
	src.EnrichRouterWithExamples(router)

	log.Println("Listening")
	log.Println(http.ListenAndServe(":"+resolvePort(), router))
}

func createServer() *http.Server {
	router := mux.NewRouter().StrictSlash(true)
	src.EnrichRouter(router)
	src.EnrichRouterWithExamples(router)

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
