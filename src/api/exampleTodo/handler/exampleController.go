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

func (h ExampleTodoHandler) CreateExampleType(ctx *fiber.Ctx) error {
	return ctx.JSON(h.service.GenerateExampleTodo())
}

func (h ExampleTodoHandler) CreateExampleIndexBody(ctx *fiber.Ctx) error {
	return ctx.JSON(h.service.CreateExampleIndexBody())
}
