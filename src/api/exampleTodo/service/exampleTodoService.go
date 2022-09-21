package exampleService

import (
	"fmt"
	"github.com/google/uuid"
	exampleDomain "github.com/strikersk/go-elastic/src/api/exampleTodo/domain"
	todoDomain "github.com/strikersk/go-elastic/src/api/todo/entity"
	elasticIndex "github.com/strikersk/go-elastic/src/elastic/core/index"
	"time"
)

type ExampleTodoService struct {
	indexBuilder elasticIndex.ElasticIndexBuilder
}

func NewExampleTodoService(indexBuilder elasticIndex.ElasticIndexBuilder) ExampleTodoService {
	return ExampleTodoService{
		indexBuilder: indexBuilder,
	}
}

func (ExampleTodoService) GenerateExampleTodo() todoDomain.Todo {
	tmpString := "Example Create Todo"
	return todoDomain.Todo{
		ID:          uuid.New().String(),
		Time:        fmt.Sprintf("%d", time.Now().Unix()),
		Name:        tmpString,
		Description: "Example Create Todo",
		Tags:        []string{"test", "testing", "qa"},
		Done:        false,
	}
}

func (r ExampleTodoService) CreateExampleIndex() elasticIndex.ElasticIndexBody {
	return r.indexBuilder.BuildIndex(exampleDomain.ExampleStruct{})
}
