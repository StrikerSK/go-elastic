package ports

import (
	"context"
	todoDomain "github.com/strikersk/go-elastic/src/api/todo/domain"
)

type ITodoRepository interface {
	SearchTodos(context.Context, string) ([]todoDomain.Todo, error)
	CreateDocument(context.Context, todoDomain.Todo) (string, error)
	ReadDocument(context.Context, string) (todoDomain.Todo, error)
	UpdateDocument(context.Context, todoDomain.Todo) error
	DeleteDocument(context.Context, string) error
}
