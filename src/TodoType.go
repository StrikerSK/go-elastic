package src

import (
	"encoding/json"
	"go-elastic/src/elastic"
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
	elasticBody := elastic.NewElasticBody(elastic.NewDefaultSettings(), *elastic.CreateMappingMap(Todo{}))

	payload, err := json.Marshal(elasticBody)
	if err != nil {
		log.Printf("Todo index initialization error: %v\n", err)
	}

	return payload
}
