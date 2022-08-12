package domain

import (
	"encoding/json"
	"github.com/strikersk/go-elastic/src/elastic/core"
	"log"
)

type Todo struct {
	ID          string `json:"id"`
	Time        string `json:"time"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

//Mapping to every type property should be made to create index
func CreateTodoIndexBody() []byte {
	elasticBody := core.NewElasticBody(core.NewDefaultSettings(), *core.CreateElasticObject(Todo{}))

	payload, err := json.Marshal(elasticBody)
	if err != nil {
		log.Printf("Todo index initialization error: %v\n", err)
	}

	return payload
}
