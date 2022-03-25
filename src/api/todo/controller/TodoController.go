package controller

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/strikersk/go-elastic/src/api/todo/entity"
	"github.com/strikersk/go-elastic/src/api/todo/repository"
	"log"
	"net/http"
	"time"
)

func CreateTodo(ctx *fiber.Ctx) error {
	todo, err := extractBody(ctx)
	if err != nil {
		return err
	}

	todo.ID = uuid.New().String()
	todo.Time = fmt.Sprintf("%d", time.Now().Unix())

	responseId, err := repository.TodoRepository.InsertDocument("", todo)
	if err != nil {
		log.Printf("Repository error: %v\n", err)
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	ctx.Status(http.StatusCreated)
	return ctx.JSON(map[string]string{"id": responseId})
}

func ReadTodo(ctx *fiber.Ctx) error {
	documentID, err := extractParam(ctx, "id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	todo, err := repository.TodoRepository.SearchDocumentByID(documentID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"data": err.Error()})
	}

	return ctx.JSON(todo)
}

func DeleteTodo(ctx *fiber.Ctx) error {
	documentID, err := extractParam(ctx, "id")
	if err != nil {
		return err
	}

	if err = repository.TodoRepository.DeleteDocument(documentID); err != nil {
		log.Printf("Delete error: %v\n", err)
	}

	return ctx.SendStatus(http.StatusOK)
}

func UpdateTodo(ctx *fiber.Ctx) error {
	todo, err := extractBody(ctx)
	if err != nil {
		return err
	}

	documentID, err := extractParam(ctx, "id")
	if err != nil {
		return err
	}

	_, err = repository.TodoRepository.InsertDocument(documentID, todo)
	if err != nil {
		log.Printf("Repository error: %v\n", err)
		return err
	}

	return ctx.SendStatus(http.StatusOK)
}

func SearchTodo(ctx *fiber.Ctx) error {
	query := struct {
		Query []string `query:"query"`
	}{}

	_ = ctx.QueryParser(&query)
	todos, _ := repository.TodoRepository.GetByStringQuery(query.Query)
	return ctx.JSON(todos)
}

func extractParam(ctx *fiber.Ctx, param string) (string, error) {
	documentID := ctx.Params(param)
	if documentID == "" {
		return "", errors.New("id parameter cannot be empty")
	}
	return documentID, nil
}

func extractBody(ctx *fiber.Ctx) (entity.Todo, error) {
	var todo entity.Todo
	if err := ctx.BodyParser(&todo); err != nil {
		log.Printf("Body parsing error: %v\n", err)
		return entity.Todo{}, err
	}

	return todo, nil
}
