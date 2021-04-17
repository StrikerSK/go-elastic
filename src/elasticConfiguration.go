package src

import (
	"context"
	"encoding/json"
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
		elastic.SetURL(HOST_URL),
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

	exists, err := elastic.ElasticClient.IndexExists(TODOS_INDEX).Do(ctx)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v\n", err)
	} else if exists {
		fmt.Printf("The index %s already exists.", TODOS_INDEX)
		if _, err = elastic.ElasticClient.DeleteIndex(TODOS_INDEX).Do(ctx); err != nil {
			log.Fatalf("client.DeleteIndex() ERROR: %v\n", err)
		}
	}

	create, err := elastic.ElasticClient.CreateIndex(TODOS_INDEX).Body(string(CreateTodoIndexBody())).Do(ctx)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v\n", err)
	} else {
		fmt.Println("CreateIndex():", create)
	}
}

func (elastic ElasticConfiguration) createData(object CustomInterface) {
	url := HOST_URL + "/" + TODOS_INDEX + "/_doc"
	method := "POST"

	marshalledObject, _ := object.MarshalItem()
	payload := strings.NewReader(string(marshalledObject))

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
	log.Print("Custom object created successfully!")

	m := make(map[string]string)
	_ = json.Unmarshal(body, &m)

	fmt.Println(m["_id"])
	fmt.Println(string(body))
}

func (elastic ElasticConfiguration) getTodo(todoID string) (todo Todo) {

	url := HOST_URL + "/" + TODOS_INDEX + "/_doc/" + todoID
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
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

	m := make(map[string]interface{})
	_ = json.Unmarshal(body, &m)

	todo.ResolveMap(m["_source"].(map[string]interface{}))
	return
}
