package exampleService

import (
	"fmt"
	"github.com/google/uuid"
	exampleDomain "github.com/strikersk/go-elastic/src/api/exampleTodo/domain"
	todoDomain "github.com/strikersk/go-elastic/src/api/todo/entity"
	elasticCore "github.com/strikersk/go-elastic/src/elastic/core"
	"time"
)

type ExampleTodoService struct {
	indexBuilder elasticCore.ElasticIndexBuilder
}

func NewExampleTodoService(indexBuilder elasticCore.ElasticIndexBuilder) ExampleTodoService {
	return ExampleTodoService{
		indexBuilder: indexBuilder,
	}
}

func (ExampleTodoService) GenerateExampleTodo() todoDomain.Todo {
	return todoDomain.Todo{
		ID:          uuid.New().String(),
		Time:        fmt.Sprintf("%d", time.Now().Unix()),
		Name:        "Example Create Todo",
		Description: "Example Create Todo",
		Done:        false,
	}
}

func (r ExampleTodoService) CreateExampleIndexBody() elasticCore.ElasticBody {
	return r.indexBuilder.BuildIndex(exampleDomain.ExampleStruct{})
}
