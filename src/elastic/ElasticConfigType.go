package elastic

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/strikersk/go-elastic/src/elastic/core"
	"log"
	"os"
	"time"
)

type ElasticConfiguration struct {
	ElasticClient *elastic.Client
	Context       context.Context
	IndexBuilder  core.ElasticIndexBuilder
}

func NewElasticConfiguration(indexBuilder core.ElasticIndexBuilder) ElasticConfiguration {
	log.Println("ElasticSearch initialization")

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(os.Getenv("ELASTIC_URL")),
		elastic.SetHealthcheckInterval(5*time.Second),
	)

	if err != nil {
		log.Printf("ElasticSearch initialization error: %s\n", err)
		os.Exit(1)
	}

	return ElasticConfiguration{
		ElasticClient: client,
		Context:       context.Background(),
		IndexBuilder:  indexBuilder,
	}
}

func (ec ElasticConfiguration) InitializeIndex(indexName string, inputStruct interface{}) {
	exists, err := ec.indexExists(indexName)
	if err != nil {
		return
	}

	if exists {
		log.Printf("Index [%s] Initialize: index already exists\n", indexName)
		if err = ec.deleteIndex(indexName); err != nil {
			return
		}
	}

	txtValue, _ := json.Marshal(ec.IndexBuilder.BuildIndex(inputStruct))
	if err = ec.createIndex(indexName, txtValue); err != nil {
		return
	}

	log.Printf("Index [%s] recreated\n", indexName)
}

func (ec ElasticConfiguration) createIndex(indexName string, indexBody []byte) error {
	if _, err := ec.ElasticClient.CreateIndex(indexName).Body(string(indexBody)).Do(ec.Context); err != nil {
		log.Printf("Index [%s] Create: %v\n", indexName, err)
		return err
	}

	log.Printf("Index [%s] Create: success\n", indexName)
	return nil
}

func (ec ElasticConfiguration) indexExists(indexName string) (bool, error) {
	exists, err := ec.ElasticClient.IndexExists(indexName).Do(ec.Context)
	if err != nil {
		log.Printf("Index [%s] existance check: %v\n", indexName, err)
		return exists, err
	} else {
		return exists, nil
	}
}

func (ec ElasticConfiguration) deleteIndex(indexName string) error {
	if _, err := ec.ElasticClient.DeleteIndex(indexName).Do(ec.Context); err != nil {
		log.Printf("Index [%s] delete: %v\n", indexName, err)
		return err
	}

	return nil
}
