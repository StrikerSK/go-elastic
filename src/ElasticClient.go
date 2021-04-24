package src

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"time"
)

var ESConfiguration = InitializeElasticSearchClient()

func InitializeElasticSearchClient() (configuration ElasticConfiguration) {
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

func (receiver ElasticConfiguration) searchTodos(searchedID string) (interface{}, error) {
	q := elastic.NewBoolQuery()
	q.Must(elastic.NewIdsQuery(searchedID))

	searchResult, err := receiver.ElasticClient.Search().
		Index(TodosIndex).
		Type("_doc").
		Query(q).
		TrackScores(false).
		Do(receiver.Context)

	if err != nil {
		log.Printf("searchTodo error %s\n", err)
		return "", err
	}

	//var todos []Todo

	for _, hit := range searchResult.Hits.Hits {
		todo, err := hit.Source.MarshalJSON()
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		fmt.Println(todo)
	}

	return nil, nil
}

func (receiver ElasticConfiguration) addTodo(input CustomInterface, indexName string) (string, error) {

	dataJSON, err := input.MarshalItem()
	if err != nil {
		log.Printf("addTodo error %s\n", err)
		return "", err
	}

	contentBody := string(dataJSON)
	replyCustom, err := receiver.ElasticClient.Index().
		Index(indexName).
		Type("_doc").
		BodyJson(contentBody).
		Do(receiver.Context)

	if err != nil {
		log.Printf("addTodo error %s\n", err)
		return "", err
	}

	return replyCustom.Id, nil
}
