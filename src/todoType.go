package src

import (
	"encoding/json"
	"log"
)

type CustomInterface interface {
	MarshalItem() ([]byte, error)
}

type Todo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"isDone"`
}

func (todo Todo) MarshalItem() ([]byte, error) {
	dataJSON, err := json.Marshal(todo)
	return dataJSON, err
}

func CreateTodoIndexBody() []byte {
	var propertyMap = make(map[string]property, 4)
	propertyMap["name"] = property{
		Type: "text",
	}

	propertyMap["description"] = property{
		Type: "text",
	}

	propertyMap["isDone"] = property{
		Type: "boolean",
	}

	//propertyMap["type"] = property{
	//	Type: "text",
	//}

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
