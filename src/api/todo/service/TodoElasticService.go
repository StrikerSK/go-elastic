package service

import (
	"fmt"
	"github.com/google/uuid"
	todoDomain "github.com/strikersk/go-elastic/src/api/todo/domain"
	todoPorts "github.com/strikersk/go-elastic/src/api/todo/ports"
	"time"
)

type TodoElasticService struct {
	repository todoPorts.ITodoRepository
}

func NewTodoElasticService(repository todoPorts.ITodoRepository) TodoElasticService {
	return TodoElasticService{
		repository: repository,
	}
}

func (s TodoElasticService) FindTodo(id string) (todoDomain.Todo, error) {
	return s.repository.FindTodo(id)
}

func (s TodoElasticService) FindTodos() ([]todoDomain.Todo, error) {
	return []todoDomain.Todo{}, nil
}

func (s TodoElasticService) CreateTodo(todo todoDomain.Todo) (string, error) {
	todo.ID = uuid.New().String()
	todo.Time = fmt.Sprintf("%d", time.Now().Unix())
	return s.repository.CreateDocument(todo)
}

func (s TodoElasticService) UpdateTodo(id string, todo todoDomain.Todo) error {
	todo.ID = id
	return s.repository.UpdateDocument(todo)
}

func (s TodoElasticService) DeleteTodo(id string) error {
	return s.repository.DeleteDocument(id)
}

func (s TodoElasticService) SearchTodos(queries []string) ([]todoDomain.Todo, error) {
	return s.repository.SearchTodos(queries)
}
