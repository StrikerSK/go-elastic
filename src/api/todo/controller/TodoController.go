package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	todoDomain "github.com/strikersk/go-elastic/src/api/todo/domain"
	todoPorts "github.com/strikersk/go-elastic/src/api/todo/ports"
	"log"
	"net/http"
)

type TodoHandler struct {
	service todoPorts.ITodoService
}

func NewTodoHandler(service todoPorts.ITodoService) TodoHandler {
	return TodoHandler{
		service: service,
	}
}

func (h TodoHandler) EnrichRouter(apiPath fiber.Router) {
	todoPath := apiPath.Group("/todo")
	todoPath.Post("", h.createTodo)
	todoPath.Put("/:id", h.updateTodo)
	todoPath.Delete("/:id", h.deleteTodo)
	todoPath.Get("/search", h.searchTodo)
	todoPath.Get("/:id", h.readTodo)
}

func (h TodoHandler) createTodo(ctx *fiber.Ctx) error {
	todo, err := h.extractBody(ctx)
	if err != nil {
		return err
	}

	responseId, err := h.service.CreateTodo(todo)
	if err != nil {
		log.Printf("Repository error: %v\n", err)
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	ctx.Status(http.StatusCreated)
	return ctx.JSON(map[string]string{"id": responseId})
}

func (h TodoHandler) readTodo(ctx *fiber.Ctx) error {
	documentID, err := h.extractParam(ctx, "id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	todo, err := h.service.FindTodo(documentID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"data": err.Error()})
	}

	return ctx.JSON(todo)
}

func (h TodoHandler) deleteTodo(ctx *fiber.Ctx) error {
	documentID, err := h.extractParam(ctx, "id")
	if err != nil {
		return err
	}

	if err = h.service.DeleteTodo(documentID); err != nil {
		log.Printf("Delete error: %v\n", err)
	}

	return ctx.SendStatus(http.StatusOK)
}

func (h TodoHandler) updateTodo(ctx *fiber.Ctx) error {
	todo, err := h.extractBody(ctx)
	if err != nil {
		return err
	}

	documentID, err := h.extractParam(ctx, "id")
	if err != nil {
		return err
	}

	if err = h.service.UpdateTodo(documentID, todo); err != nil {
		log.Printf("Repository error: %v\n", err)
		return err
	}

	return ctx.SendStatus(http.StatusOK)
}

func (h TodoHandler) searchTodo(ctx *fiber.Ctx) error {
	query := struct {
		Query []string `query:"query"`
	}{}

	_ = ctx.QueryParser(&query)
	todos, _ := h.service.SearchTodos(query.Query)
	return ctx.JSON(todos)
}

func (h TodoHandler) extractParam(ctx *fiber.Ctx, param string) (string, error) {
	documentID := ctx.Params(param)
	if documentID == "" {
		return "", errors.New("id parameter cannot be empty")
	}
	return documentID, nil
}

func (h TodoHandler) extractBody(ctx *fiber.Ctx) (todoDomain.Todo, error) {
	var todo todoDomain.Todo
	if err := ctx.BodyParser(&todo); err != nil {
		log.Printf("Body parsing error: %v\n", err)
		return todoDomain.Todo{}, err
	}

	return todo, nil
}
