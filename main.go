package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-elastic/src/api/exampleType"
	"github.com/strikersk/go-elastic/src/api/todo/controller"
	"github.com/strikersk/go-elastic/src/elastic"
	"log"
	"os"
)

func init() {
	elastic.GetElasticInstance()
}

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})

	apiPath := app.Group("/api")

	examplePath := apiPath.Group("/examples")
	examplePath.Get("/index", exampleType.CreateExampleType)
	examplePath.Get("/type", exampleType.CreateExampleIndexBody)

	todoPath := app.Group("/todo")
	todoPath.Post("", controller.CreateTodo)
	todoPath.Put("/:id", controller.UpdateTodo)
	todoPath.Delete("/:id", controller.DeleteTodo)
	todoPath.Get("/search", controller.SearchTodo)
	todoPath.Get("/:id", controller.ReadTodo)

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
