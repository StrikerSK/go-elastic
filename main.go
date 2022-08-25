package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	exampleHandler "github.com/strikersk/go-elastic/src/api/exampleTodo/handler"
	exampleService "github.com/strikersk/go-elastic/src/api/exampleTodo/service"
	todoController "github.com/strikersk/go-elastic/src/api/todo/controller"
	todoRepository "github.com/strikersk/go-elastic/src/api/todo/repository"
	elasticConfig "github.com/strikersk/go-elastic/src/elastic/config"
	elasticCore "github.com/strikersk/go-elastic/src/elastic/core"
	"log"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})

	mappingFactory := elasticCore.NewElasticMappingFactory()
	indexBuilder := elasticCore.NewElasticIndexBuilder(mappingFactory)
	elasticConfiguration := elasticConfig.NewElasticConfiguration(indexBuilder)

	exSrv := exampleService.NewExampleTodoService(indexBuilder)
	exHdl := exampleHandler.NewExampleTodoHandler(exSrv)

	apiPath := app.Group("/api")

	examplePath := apiPath.Group("/examples")
	examplePath.Get("/index", exHdl.CreateExampleType)
	examplePath.Get("/type", exHdl.CreateExampleIndexBody)

	elasticRepo := todoRepository.NewElasticRepository(elasticConfiguration)
	handler := todoController.NewTodoHandler(elasticRepo)

	todoPath := apiPath.Group("/todo")
	todoPath.Post("", handler.CreateTodo)
	todoPath.Put("/:id", handler.UpdateTodo)
	todoPath.Delete("/:id", handler.DeleteTodo)
	todoPath.Get("/search", handler.SearchTodo)
	todoPath.Get("/:id", handler.ReadTodo)

	log.Fatal(app.Listen(resolvePort()))
}

func resolvePort() (port string) {
	port = os.Getenv("PORT")

	if port == "" {
		log.Printf("Default PORT value used\n")
		port = "5000"
	}

	return fmt.Sprintf(":%s", port)
}
