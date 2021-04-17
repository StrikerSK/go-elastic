package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-elastic/src"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	src.ESConfiguration.CreateElasticIndex()

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "5000"
	}

	myRouter := mux.NewRouter().StrictSlash(true)
	src.EnrichRouter(myRouter)
	fmt.Println("Listening")

	fmt.Println(http.ListenAndServe(":"+port, myRouter))
}
