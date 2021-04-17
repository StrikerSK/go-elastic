package src

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var ESConfiguration = InitializeElasticSearchClient()

func InitializeElasticSearchClient() ElasticConfiguration {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(5*time.Second),
	)

	if err != nil {
		panic(err)
	}

	log.Print("ElasticSearch initialized...")
	return ElasticConfiguration{
		ElasticClient: client,
	}
}

type ElasticConfiguration struct {
	ElasticClient *elastic.Client
}

func (elastic ElasticConfiguration) CreateElasticIndex() {
	ctx := context.Background()

	exists, err := elastic.ElasticClient.IndexExists(todosIndex).Do(ctx)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v\n", err)
	} else if exists {
		fmt.Printf("The index %s already exists.", todosIndex)
		if _, err = elastic.ElasticClient.DeleteIndex(todosIndex).Do(ctx); err != nil {
			log.Fatalf("client.DeleteIndex() ERROR: %v\n", err)
		}
	}

	create, err := elastic.ElasticClient.CreateIndex(todosIndex).Body(string(CreateTodoIndexBody())).Do(ctx)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v\n", err)
	} else {
		fmt.Println("CreateIndex():", create)
	}
}

func (elastic ElasticConfiguration) createData(todo Todo) {
	url := "http://localhost:9200/custom_todos/_doc"
	method := "POST"

	payload := strings.NewReader(string(todo.marshalTodo()))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
