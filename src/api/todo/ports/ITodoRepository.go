package ports

import (
	todoDomain "github.com/strikersk/go-elastic/src/api/todo/domain"
)

type ITodoRepository interface {
	FindTodo(string) (todoDomain.Todo, error)
	SearchTodos(string) ([]todoDomain.Todo, error)
	CreateDocument(todoDomain.Todo) (string, error)
	UpdateDocument(todoDomain.Todo) error
	DeleteDocument(string) error
}
