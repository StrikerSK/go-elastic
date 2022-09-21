package repository

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	todoDomain "github.com/strikersk/go-elastic/src/api/todo/domain"
	elasticConfig "github.com/strikersk/go-elastic/src/elastic/config"
	"log"
	"reflect"
)

type TodoElasticRepository struct {
	client    *elastic.Client
	context   context.Context
	indexName string
}

func NewElasticRepository(config elasticConfig.ElasticConfiguration) *TodoElasticRepository {
	indexName := "todos"
	err := config.InitializeIndex(indexName, todoDomain.Todo{})
	if err != nil {
		panic(err)
	}

	return &TodoElasticRepository{
		client:    config.ElasticClient,
		context:   config.Context,
		indexName: indexName,
	}
}

func (r TodoElasticRepository) DeleteDocument(documentID string) (err error) {
	_, err = r.client.Delete().Index(r.indexName).Id(documentID).Do(r.context)
	return
}

func (r TodoElasticRepository) CreateDocument(document todoDomain.Todo) (string, error) {
	return r.insertDocument("", document)
}

func (r TodoElasticRepository) UpdateDocument(document todoDomain.Todo) error {
	_, err := r.insertDocument(document.ID, document)
	return err
}

func (r TodoElasticRepository) insertDocument(id string, document todoDomain.Todo) (string, error) {
	data, err := json.Marshal(document)
	if err != nil {
		log.Printf("Marshalling document error: %v\n", err)
		return "", err
	}

	response, err := r.client.Index().Index(r.indexName).Id(id).BodyJson(string(data)).Do(r.context)
	if err != nil {
		log.Printf("Insert to index error: %v\n", err)
		return "", err
	}

	log.Println("Document inserted to index")
	return response.Id, nil
}

func (r TodoElasticRepository) FindTodo(documentID string) (output todoDomain.Todo, err error) {
	searchResult, err := r.client.Get().Index(r.indexName).Id(documentID).Do(r.context)
	if err != nil {
		log.Printf("Searching error: %v\n", err)
		return todoDomain.Todo{}, err
	}

	resolvedStructure, err := searchResult.Source.MarshalJSON()
	if err != nil {
		log.Printf("Marshalling error: %v\n", err)
		return todoDomain.Todo{}, err
	}

	if err = json.Unmarshal(resolvedStructure, &output); err != nil {
		log.Printf("Unmarshalling error: %v\n", err)
		return todoDomain.Todo{}, err
	}

	log.Println("Search successful")
	return
}

func (r TodoElasticRepository) SearchTodos(stringQuery []string) ([]todoDomain.Todo, error) {
	output := make([]todoDomain.Todo, 0)

	searchService := r.client.Search().Index(r.indexName) // search in index "todos"

	for _, query := range stringQuery {
		searchService = searchService.Query(elastic.NewQueryStringQuery(query))
	}

	searchResult, err := searchService.
		//Sort("name", true). // sort by "user" field, ascending
		//From(0).Size(2). // take documents 0-9
		Pretty(true).
		Do(r.context)

	if err != nil {
		return nil, err
	}

	for _, item := range searchResult.Each(reflect.TypeOf(todoDomain.Todo{})) {
		todo := item.(todoDomain.Todo)
		output = append(output, todo)
	}

	return output, nil
}
