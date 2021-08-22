package src

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"sync"
	"time"
)

type ElasticConfiguration struct {
	ElasticClient *elastic.Client
	Context       context.Context
}

var elasticLock = &sync.Mutex{}
var elasticConfiguration *ElasticConfiguration

func GetElasticInstance() *ElasticConfiguration {
	//To prevent expensive lock operations
	//This means that the cacheConnection field is already populated
	if elasticConfiguration == nil {
		elasticLock.Lock()
		defer elasticLock.Unlock()

		//Only one goroutine can create the singleton instance.
		if elasticConfiguration == nil {
			var configuration ElasticConfiguration
			log.Println("Creating ElasticSearch instance")
			client, err := elastic.NewClient(
				elastic.SetSniff(false),
				elastic.SetURL(os.Getenv("ELASTIC_URL")),
				elastic.SetHealthcheckInterval(5*time.Second),
			)

			if err != nil {
				log.Printf("ElasticSearch Initialization error: %s\n", err)
				os.Exit(1)
			}

			configuration.ElasticClient = client
			configuration.Context = context.Background()
			elasticConfiguration = &configuration
			log.Println("ElasticSearch initialized...")
		} else {
			log.Println("ElasticSearch instance already created!")
		}
	} else {
		//log.Println("Application Cache instance already created!")
	}

	return elasticConfiguration
}

func (ec ElasticConfiguration) CreateIndex(indexName string, indexBody []byte) {
	exists, err := ec.ElasticClient.IndexExists(indexName).Do(ec.Context)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v\n", err)
	} else if exists {
		fmt.Printf("Index [%s] already exists.\n", indexName)
		if _, err = ec.ElasticClient.DeleteIndex(indexName).Do(ec.Context); err != nil {
			log.Fatalf("client.DeleteIndex() ERROR: %v\n", err)
		}
	}

	_, err = ec.ElasticClient.CreateIndex(indexName).Body(string(indexBody)).Do(ec.Context)
	if err != nil {
		log.Printf("Create Index [%s] error: %v\n", indexName, err)
		return
	} else {
		fmt.Printf("Index [%s] created\n", indexName)
		return
	}
}

func (ec ElasticConfiguration) searchTodos(indexName string, searchedID string, targetClass CustomInterface) error {
	searchResult, err := ec.ElasticClient.
		Get().
		Index(indexName).
		Id(searchedID).
		Do(ec.Context)

	//TODO Create solution to transfer status code
	//This might be always not found
	if err != nil {
		log.Printf("searchTodo error %s\n", err)
		return err
	}

	test, err := searchResult.Source.MarshalJSON()
	if err != nil {
		log.Printf("searchTodo error %s\n", err)
		return err
	}

	if err = targetClass.UnmarshalItem(test); err != nil {
		log.Printf("searchTodo error %s\n", err)
		return err
	}

	return nil
}

func (ec ElasticConfiguration) insertToIndex(itemId string, input CustomInterface, indexName string) (string, error) {

	dataJSON, err := input.MarshalItem()
	if err != nil {
		log.Printf("Insert to Index [%s] error: %s\n", indexName, err)
		return "", err
	}

	contentBody := string(dataJSON)
	replyCustom, err := ec.ElasticClient.Index().
		Index(indexName).
		Id(itemId).
		BodyJson(contentBody).
		Do(ec.Context)

	if err != nil {
		log.Printf("Insert to Index [%s] error: %s\n", indexName, err)
		return "", err
	}

	return replyCustom.Id, nil
}

func (ec ElasticConfiguration) deleteItem(indexName string, searchedID string) {
	_, err := ec.ElasticClient.
		Delete().
		Index(indexName).
		Id(searchedID).
		Do(ec.Context)

	if err != nil {
		log.Printf("Delete Index error: %s", err)
		return
	}
}
