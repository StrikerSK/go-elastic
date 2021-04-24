package src

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"time"
)

var ESConfiguration = InitializeElasticSearchClient()

func InitializeElasticSearchClient() ElasticConfiguration {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(HostUrl),
		elastic.SetHealthcheckInterval(5*time.Second),
	)

	if err != nil {
		panic(err)
	}

	log.Println("ElasticSearch initialized...")
	return ElasticConfiguration{
		ElasticClient: client,
		Context:       context.Background(),
	}
}

type ElasticConfiguration struct {
	ElasticClient *elastic.Client
	Context       context.Context
}

func (elastic ElasticConfiguration) CreateElasticIndex() {
	exists, err := elastic.ElasticClient.IndexExists(TodosIndex).Do(elastic.Context)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v\n", err)
	} else if exists {
		fmt.Printf("The index %s already exists.\n", TodosIndex)
		if _, err = elastic.ElasticClient.DeleteIndex(TodosIndex).Do(elastic.Context); err != nil {
			log.Fatalf("client.DeleteIndex() ERROR: %v\n", err)
		}
	}

	create, err := elastic.ElasticClient.CreateIndex(TodosIndex).Body(string(CreateTodoIndexBody())).Do(elastic.Context)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v\n", err)
	} else {
		fmt.Println("CreateIndex():", create)
	}
}
