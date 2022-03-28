package repository

import (
	"github.com/strikersk/go-elastic/src/api/todo/entity"
)

var TodoRepository ITodoRepository

func SetTodoRepository(input ITodoRepository) {
	TodoRepository = input
}

type ITodoRepository interface {
	GetByID(string) (entity.Todo, error)
	SearchByStringQuery([]string) ([]entity.Todo, error)
	InsertDocument(string, entity.Todo) (string, error)
	DeleteDocument(string) error
}
