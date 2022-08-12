package controller

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/strikersk/go-elastic/src/api/todo/entity"
	"github.com/strikersk/go-elastic/src/api/todo/ports"
	"log"
	"net/http"
	"time"
)

type TodoHandler struct {
	service ports.ITodoRepository
}

func NewTodoHandler(service ports.ITodoRepository) TodoHandler {
	return TodoHandler{
		service: service,
	}
}

func (h TodoHandler) CreateTodo(ctx *fiber.Ctx) error {
	todo, err := extractBody(ctx)
	if err != nil {
		return err
	}

	todo.ID = uuid.New().String()
	todo.Time = fmt.Sprintf("%d", time.Now().Unix())

	responseId, err := h.service.InsertDocument("", todo)
	if err != nil {
		log.Printf("Repository error: %v\n", err)
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	ctx.Status(http.StatusCreated)
	return ctx.JSON(map[string]string{"id": responseId})
}

func (h TodoHandler) ReadTodo(ctx *fiber.Ctx) error {
	documentID, err := extractParam(ctx, "id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	todo, err := h.service.SearchByID(documentID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"data": err.Error()})
	}

	return ctx.JSON(todo)
}

func (h TodoHandler) DeleteTodo(ctx *fiber.Ctx) error {
	documentID, err := extractParam(ctx, "id")
	if err != nil {
		return err
	}

	if err = h.service.DeleteDocument(documentID); err != nil {
		log.Printf("Delete error: %v\n", err)
	}

	return ctx.SendStatus(http.StatusOK)
}

func (h TodoHandler) UpdateTodo(ctx *fiber.Ctx) error {
	todo, err := extractBody(ctx)
	if err != nil {
		return err
	}

	documentID, err := extractParam(ctx, "id")
	if err != nil {
		return err
	}

	_, err = h.service.InsertDocument(documentID, todo)
	if err != nil {
		log.Printf("Repository error: %v\n", err)
		return err
	}

	return ctx.SendStatus(http.StatusOK)
}

func (h TodoHandler) SearchTodo(ctx *fiber.Ctx) error {
	query := struct {
		Query []string `query:"query"`
	}{}

	_ = ctx.QueryParser(&query)
	todos, _ := h.service.SearchByStringQuery(query.Query)
	return ctx.JSON(todos)
}

func extractParam(ctx *fiber.Ctx, param string) (string, error) {
	documentID := ctx.Params(param)
	if documentID == "" {
		return "", errors.New("id parameter cannot be empty")
	}
	return documentID, nil
}

func extractBody(ctx *fiber.Ctx) (domain.Todo, error) {
	var todo domain.Todo
	if err := ctx.BodyParser(&todo); err != nil {
		log.Printf("Body parsing error: %v\n", err)
		return domain.Todo{}, err
	}

	return todo, nil
}
