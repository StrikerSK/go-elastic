package exampleType

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/strikersk/go-elastic/src/api/todo/entity"
	"github.com/strikersk/go-elastic/src/elastic/body"
	"github.com/strikersk/go-elastic/src/response"
	"net/http"
	"time"
)

func CreateExampleType(ctx *fiber.Ctx) error {
	customTodo := entity.Todo{
		ID:          uuid.New().String(),
		Time:        fmt.Sprintf("%d", time.Now().Unix()),
		Name:        "Example Create Todo",
		Description: "Example Create Todo",
		Done:        false,
	}

	res := response.NewRequestResponse(http.StatusOK, customTodo)
	return ctx.JSON(res)
}

func CreateExampleIndexBody(ctx *fiber.Ctx) error {
	res := response.NewRequestResponse(http.StatusOK, body.NewDefaultElasticBody(*body.CreateMappingMap(exampleStruct{})))
	return ctx.JSON(res)
}
