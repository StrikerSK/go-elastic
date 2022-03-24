package main

import (
	"github.com/gorilla/mux"
	"github.com/strikersk/go-elastic/src/api/exampleType"
	"github.com/strikersk/go-elastic/src/api/todo/controller"
	"github.com/strikersk/go-elastic/src/api/todo/entity"
	"github.com/strikersk/go-elastic/src/api/todo/repository"
	"github.com/strikersk/go-elastic/src/elastic"
	"log"
	"net/http"
	"os"
)

func init() {
	elastic.GetElasticInstance().InitializeIndex(repository.TodoIndex, entity.CreateTodoIndexBody())
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	controller.EnrichRouter(router)
	exampleType.EnrichRouterWithExamples(router)

	log.Println("Listening")
	log.Println(http.ListenAndServe(":"+resolvePort(), router))
}

func createServer() *http.Server {
	router := mux.NewRouter().StrictSlash(true)
	controller.EnrichRouter(router)
	exampleType.EnrichRouterWithExamples(router)

	return &http.Server{
		Addr:    resolvePort(),
		Handler: router,
	}
}

func resolvePort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		log.Printf("Default PORT value used\n")
		port = "5000"
	}
	return
}
