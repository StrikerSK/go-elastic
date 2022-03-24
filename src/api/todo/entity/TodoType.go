package entity

import (
	"encoding/json"
	"github.com/strikersk/go-elastic/src/elastic/body"
	"log"
)

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
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
