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

func (todo *Todo) ResolveMap(inputMap map[string]interface{}) {
	todo.Name = inputMap["name"].(string)
	todo.Description = inputMap["description"].(string)
	todo.Done = inputMap["done"].(bool)
	return
}

func CreateTodoIndexBody() []byte {
	var propertyMap = make(map[string]property, 4)
	propertyMap["name"] = property{
		Type: "text",
	}

	propertyMap["description"] = property{
		Type: "text",
	}

	propertyMap["done"] = property{
		Type: "boolean",
	}

	elasticBody := elasticBody{
		Settings: settings{
			NumberOfShards:   1,
			NumberOfReplicas: 1,
		},
		Mappings: mappings{
			Properties: propertyMap,
		},
	}

	payload, err := json.Marshal(elasticBody)
	if err != nil {
		log.Fatal(err)
	}

	return payload
}
