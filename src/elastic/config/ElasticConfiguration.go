package elasticConfig

import (
	"context"
	"github.com/olivere/elastic/v7"
	elasticIndex "github.com/strikersk/go-elastic/src/elastic/core/index"
	"log"
	"os"
	"time"
)

type ElasticConfiguration struct {
	ElasticClient *elastic.Client
	Context       context.Context
	IndexBuilder  elasticIndex.ElasticIndexBuilder
}

func NewElasticConfiguration(indexBuilder elasticIndex.ElasticIndexBuilder) ElasticConfiguration {
	log.Println("ElasticSearch initialization")

	elasticUrl := os.Getenv("ELASTIC_URL")
	if elasticUrl == "" {
		elasticUrl = "http://localhost:9200"
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(elasticUrl),
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

func (ec ElasticConfiguration) InitializeIndex(indexName string, inputStruct interface{}) error {
	exists, err := ec.indexExists(indexName)
	if err != nil {
		return err
	}

	if exists {
		log.Printf("Index [%s] recreating!\n", indexName)
		if err = ec.deleteIndex(indexName); err != nil {
			return err
		}
	} else {
		log.Printf("Index [%s] creating!\n", indexName)
	}

	indexData, err := ec.IndexBuilder.BuildAndMarshallIndex(inputStruct)
	if err != nil {
		return err
	}

	if err = ec.createIndex(indexName, indexData); err != nil {
		return err
	}

	return nil
}

func (ec ElasticConfiguration) createIndex(indexName string, indexBody []byte) error {
	if _, err := ec.ElasticClient.CreateIndex(indexName).Body(string(indexBody)).Do(ec.Context); err != nil {
		log.Printf("Index [%s] Create: %v\n", indexName, err)
		return err
	}

	log.Printf("Index [%s] created successfully!\n", indexName)
	return nil
}

func (ec ElasticConfiguration) indexExists(indexName string) (bool, error) {
	exists, err := ec.ElasticClient.IndexExists(indexName).Do(ec.Context)
	if err != nil {
		log.Printf("Index [%s] exist error: %v\n", indexName, err)
		return false, err
	}

	return exists, nil
}

func (ec ElasticConfiguration) deleteIndex(indexName string) error {
	if _, err := ec.ElasticClient.DeleteIndex(indexName).Do(ec.Context); err != nil {
		log.Printf("Index [%s] delete error: %v\n", indexName, err)
		return err
	}

	return nil
}
