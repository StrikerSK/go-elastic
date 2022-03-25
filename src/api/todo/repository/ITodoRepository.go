package repository

import (
	"github.com/strikersk/go-elastic/src/api/todo/entity"
)

type ITodoRepository interface {
	SearchDocumentByID(string) (entity.Todo, error)
	GetByStringQuery([]string) ([]entity.Todo, error)
	InsertDocument(string, entity.Todo) (string, error)
	DeleteDocument(string) error
}

var TodoRepository ITodoRepository

func SetTodoRepository(input ITodoRepository) {
	TodoRepository = input
}
