package exampleHandler

import (
	"github.com/gofiber/fiber/v2"
	exampleService "github.com/strikersk/go-elastic/src/api/exampleTodo/service"
)

type ExampleTodoHandler struct {
	service exampleService.ExampleTodoService
}

func NewExampleTodoHandler(service exampleService.ExampleTodoService) ExampleTodoHandler {
	return ExampleTodoHandler{
		service: service,
	}
}

func (h ExampleTodoHandler) EnrichHandler(api fiber.Router) {
	examplePath := api.Group("/examples")
	examplePath.Get("/index", h.CreateExampleIndex)
	examplePath.Get("/type", h.CreateExampleStruct)
}

func (h ExampleTodoHandler) CreateExampleStruct(ctx *fiber.Ctx) error {
	return ctx.JSON(h.service.GenerateExampleStruct())
}

func (h ExampleTodoHandler) CreateExampleIndex(ctx *fiber.Ctx) error {
	return ctx.JSON(h.service.CreateExampleIndex())
}
