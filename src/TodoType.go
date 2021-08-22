package src

import (
	"encoding/json"
	"log"
)

type CustomInterface interface {
	MarshalItem() ([]byte, error)
	UnmarshalItem([]byte) error
}

type Todo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (todo Todo) MarshalItem() ([]byte, error) {
	dataJSON, err := json.Marshal(todo)
	return dataJSON, err
}

func (todo *Todo) UnmarshalItem(input []byte) error {
	if err := json.Unmarshal(input, todo); err != nil {
		return err
	}
	return nil
}

//Mapping to every type property should be made to create index
func CreateTodoIndexBody() []byte {
	elasticBody := elasticBody{
		Settings: settings{
			NumberOfShards:   1,
			NumberOfReplicas: 1,
		},
		Mappings: *CreateMapping(Todo{}),
	}

	payload, err := json.Marshal(elasticBody)
	if err != nil {
		log.Printf("Todo index initialization error: %v\n", err)
	}

	return payload
}
