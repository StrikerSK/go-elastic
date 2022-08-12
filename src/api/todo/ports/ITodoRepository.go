package ports

import (
	"github.com/strikersk/go-elastic/src/api/todo/entity"
)

type ITodoRepository interface {
	SearchByID(string) (domain.Todo, error)
	SearchByStringQuery([]string) ([]domain.Todo, error)
	InsertDocument(string, domain.Todo) (string, error)
	DeleteDocument(string) error
}
