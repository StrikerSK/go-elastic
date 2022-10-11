package service

import (
	"context"
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

func (s TodoElasticService) FindTodo(ctx context.Context, id string) (todoDomain.Todo, error) {
	return s.repository.ReadDocument(ctx, id)
}

func (s TodoElasticService) FindTodos(ctx context.Context) ([]todoDomain.Todo, error) {
	return s.repository.SearchTodos(ctx, "")
}

func (s TodoElasticService) CreateTodo(ctx context.Context, todo todoDomain.Todo) (string, error) {
	todo.ID = uuid.New().String()
	todo.CreatedAt = fmt.Sprintf("%d", time.Now().Unix())
	todo.CheckDone()
	return s.repository.CreateDocument(ctx, todo)
}

func (s TodoElasticService) UpdateTodo(ctx context.Context, id string, todo todoDomain.Todo) error {
	todo.ID = id
	todo.CheckDone()
	return s.repository.UpdateDocument(ctx, todo)
}

func (s TodoElasticService) DeleteTodo(ctx context.Context, id string) error {
	return s.repository.DeleteDocument(ctx, id)
}

func (s TodoElasticService) SearchTodos(ctx context.Context, query string) ([]todoDomain.Todo, error) {
	return s.repository.SearchTodos(ctx, query)
}
