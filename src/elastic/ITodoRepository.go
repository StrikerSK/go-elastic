package elastic

import "github.com/strikersk/go-elastic/src/types"

type ITodoRepository interface {
	SearchDocument(string, types.MarshallingInterface) error
	InsertDocument(string, types.MarshallingInterface) (string, error)
	DeleteDocument(string, string)
}
