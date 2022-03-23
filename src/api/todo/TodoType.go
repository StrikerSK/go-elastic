package todo

import (
	"encoding/json"
	"github.com/strikersk/go-elastic/src/elastic/body"
	"log"
)

type Todo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (todo *Todo) MarshalItem() ([]byte, error) {
	bs, err := json.Marshal(todo)
	return bs, err
}

func (todo *Todo) UnmarshalItem(bs []byte) error {
	if err := json.Unmarshal(bs, todo); err != nil {
		return err
	}
	return nil
}

func (todo Todo) GetIndexName() string {
	return TodosIndex
}

//Mapping to every type property should be made to create index
func CreateTodoIndexBody() []byte {
	elasticBody := body.NewElasticBody(body.NewDefaultSettings(), *body.CreateMappingMap(Todo{}))

	payload, err := json.Marshal(elasticBody)
	if err != nil {
		log.Printf("Todo index initialization error: %v\n", err)
	}

	return payload
}
