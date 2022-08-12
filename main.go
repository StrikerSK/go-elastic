package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-elastic/src/api/exampleType"
	"github.com/strikersk/go-elastic/src/api/todo/controller"
	"github.com/strikersk/go-elastic/src/api/todo/repository"
	"github.com/strikersk/go-elastic/src/elastic"
	"github.com/strikersk/go-elastic/src/elastic/core"
	"log"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})

	mappingFactory := core.NewElasticMappingFactory()
	indexBuilder := core.NewElasticIndexBuilder(mappingFactory)
	exampleHandler := exampleType.NewExampleTypeHandler(indexBuilder)

	apiPath := app.Group("/api")
	examplePath := apiPath.Group("/examples")
	examplePath.Get("/index", exampleHandler.CreateExampleType)
	examplePath.Get("/type", exampleHandler.CreateExampleIndexBody)

	elasticConfiguration := elastic.NewElasticConfiguration(indexBuilder)
	elasticRepo := repository.NewElasticRepository(elasticConfiguration)
	handler := controller.NewTodoHandler(elasticRepo)

	todoPath := app.Group("/todo")
	todoPath.Post("", handler.CreateTodo)
	todoPath.Put("/:id", handler.UpdateTodo)
	todoPath.Delete("/:id", handler.DeleteTodo)
	todoPath.Get("/search", handler.SearchTodo)
	todoPath.Get("/:id", handler.ReadTodo)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", resolvePort())))
}

func resolvePort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		log.Printf("Default PORT value used\n")
		port = "5000"
	}
	return
}
