package exampleType

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	domain "github.com/strikersk/go-elastic/src/api/todo/entity"
	"github.com/strikersk/go-elastic/src/elastic/core"
	"github.com/strikersk/go-elastic/src/response"
	"net/http"
	"time"
)

type ExampleTypeHadler struct {
	indexBuilder core.ElasticIndexBuilder
}

func NewExampleTypeHandler(indexBuilder core.ElasticIndexBuilder) ExampleTypeHadler {
	return ExampleTypeHadler{
		indexBuilder: indexBuilder,
	}
}

func (ExampleTypeHadler) CreateExampleType(ctx *fiber.Ctx) error {
	customTodo := domain.Todo{
		ID:          uuid.New().String(),
		Time:        fmt.Sprintf("%d", time.Now().Unix()),
		Name:        "Example Create Todo",
		Description: "Example Create Todo",
		Done:        false,
	}

	res := response.NewRequestResponse(http.StatusOK, customTodo)
	return ctx.JSON(res)
}

func (r ExampleTypeHadler) CreateExampleIndexBody(ctx *fiber.Ctx) error {
	newIndex := r.indexBuilder.BuildIndex(exampleStruct{})
	res := response.NewRequestResponse(http.StatusOK, newIndex)
	return ctx.JSON(res)
}
