package ports

import (
	"context"
	todoDomain "github.com/strikersk/go-elastic/src/api/todo/domain"
)

type ITodoService interface {
	FindTodo(context.Context, string) (todoDomain.Todo, error)
	FindTodos(context.Context) ([]todoDomain.Todo, error)
	SearchTodos(context.Context, string) ([]todoDomain.Todo, error)
	CreateTodo(context.Context, todoDomain.Todo) (string, error)
	UpdateTodo(context.Context, string, todoDomain.Todo) error
	DeleteTodo(context.Context, string) error
}
