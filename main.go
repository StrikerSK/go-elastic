package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	exampleHandler "github.com/strikersk/go-elastic/src/api/exampleTodo/handler"
	exampleService "github.com/strikersk/go-elastic/src/api/exampleTodo/service"
	todoController "github.com/strikersk/go-elastic/src/api/todo/controller"
	todoRepository "github.com/strikersk/go-elastic/src/api/todo/repository"
	elasticService "github.com/strikersk/go-elastic/src/api/todo/service"
	elasticConfig "github.com/strikersk/go-elastic/src/elastic/config"
	elasticIndex "github.com/strikersk/go-elastic/src/elastic/core/index"
	elasticMappings "github.com/strikersk/go-elastic/src/elastic/core/mappings"
	"log"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})

	mappingFactory := elasticMappings.NewElasticMappingFactory()
	indexBuilder := elasticIndex.NewElasticIndexBuilder(mappingFactory)
	elasticConfiguration := elasticConfig.NewElasticConfiguration(indexBuilder)

	appExampleServer := exampleService.NewExampleTodoService(indexBuilder)
	appExampleHandler := exampleHandler.NewExampleTodoHandler(appExampleServer)

	elasticTodoRepository := todoRepository.NewElasticRepository(elasticConfiguration)
	elasticTodoService := elasticService.NewTodoElasticService(elasticTodoRepository)
	elasticTodoHandler := todoController.NewTodoHandler(elasticTodoService)

	apiPath := app.Group("/api")
	appExampleHandler.EnrichHandler(apiPath)
	elasticTodoHandler.EnrichRouter(apiPath)

	log.Fatal(app.Listen(resolvePort()))
}

func resolvePort() (port string) {
	port = os.Getenv("PORT")

	if port == "" {
		log.Printf("Default PORT value used\n")
		port = "4000"
	}

	return fmt.Sprintf(":%s", port)
}
