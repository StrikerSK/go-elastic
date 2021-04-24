package src

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"time"
)

var ESConfiguration = initializeElasticSearchClient()

func initializeElasticSearchClient() (configuration ElasticConfiguration) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(HostUrl),
		elastic.SetHealthcheckInterval(5*time.Second),
	)

	if err != nil {
		panic(err)
	}

	log.Println("ElasticSearch initialized...")

	configuration.ElasticClient = client
	configuration.Context = context.Background()
	configuration.createElasticIndex()

	return
}

type ElasticConfiguration struct {
	ElasticClient *elastic.Client
	Context       context.Context
}

func (receiver ElasticConfiguration) createElasticIndex() {
	exists, err := receiver.ElasticClient.IndexExists(TodosIndex).Do(receiver.Context)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v\n", err)
	} else if exists {
		fmt.Printf("The index %s already exists.\n", TodosIndex)
		if _, err = receiver.ElasticClient.DeleteIndex(TodosIndex).Do(receiver.Context); err != nil {
			log.Fatalf("client.DeleteIndex() ERROR: %v\n", err)
		}
	}

	create, err := receiver.ElasticClient.CreateIndex(TodosIndex).Body(string(CreateTodoIndexBody())).Do(receiver.Context)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v\n", err)
	} else {
		fmt.Println("CreateIndex():", create)
	}
}

func (receiver ElasticConfiguration) searchTodos(indexName string, searchedID string, targetClass CustomInterface) error {
	searchResult, err := receiver.ElasticClient.
		Get().
		Index(indexName).
		Id(searchedID).
		Do(receiver.Context)

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

func (receiver ElasticConfiguration) insertToIndex(itemId string, input CustomInterface, indexName string) (string, error) {

	dataJSON, err := input.MarshalItem()
	if err != nil {
		log.Printf("insertToIndex() error %s\n", err)
		return "", err
	}

	contentBody := string(dataJSON)
	replyCustom, err := receiver.ElasticClient.Index().
		Index(indexName).
		Id(itemId).
		BodyJson(contentBody).
		Do(receiver.Context)

	if err != nil {
		log.Printf("insertToIndex() error %s\n", err)
		return "", err
	}

	return replyCustom.Id, nil
}

func (receiver ElasticConfiguration) deleteItem(indexName string, searchedID string) {
	_, err := receiver.ElasticClient.
		Delete().
		Index(indexName).
		Id(searchedID).
		Do(receiver.Context)

	if err != nil {
		log.Printf("deleteItem() error: %s", err)
	}
}
