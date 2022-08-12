package repository

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/strikersk/go-elastic/src/api/todo/entity"
	elastic2 "github.com/strikersk/go-elastic/src/elastic"
	"log"
	"reflect"
)

type TodoElasticRepository struct {
	client    *elastic.Client
	context   context.Context
	indexName string
}

func NewElasticRepository(config elastic2.ElasticConfiguration) *TodoElasticRepository {
	indexName := "todos"
	config.InitializeIndex(indexName, domain.Todo{})

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

func (r TodoElasticRepository) InsertDocument(documentID string, document domain.Todo) (string, error) {
	data, err := json.Marshal(document)
	if err != nil {
		log.Printf("Marshalling document error: %v\n", err)
		return "", err
	}

	response, err := r.client.Index().Index(r.indexName).Id(documentID).BodyJson(string(data)).Do(r.context)
	if err != nil {
		log.Printf("Insert to index error: %v\n", err)
		return "", err
	}

	log.Println("Document inserted to index")
	return response.Id, nil
}

func (r TodoElasticRepository) SearchByID(documentID string) (output domain.Todo, err error) {
	searchResult, err := r.client.Get().Index(r.indexName).Id(documentID).Do(r.context)
	if err != nil {
		log.Printf("Searching error: %v\n", err)
		return domain.Todo{}, err
	}

	resolvedStructure, err := searchResult.Source.MarshalJSON()
	if err != nil {
		log.Printf("Marshalling error: %v\n", err)
		return domain.Todo{}, err
	}

	if err = json.Unmarshal(resolvedStructure, &output); err != nil {
		log.Printf("Unmarshalling error: %v\n", err)
		return domain.Todo{}, err
	}

	log.Println("Search successful")
	return
}

func (r TodoElasticRepository) SearchByStringQuery(stringQuery []string) ([]domain.Todo, error) {
	output := make([]domain.Todo, 0)

	searchService := r.client.Search().Index(r.indexName) // search in index "todos"

	for _, query := range stringQuery {
		searchService = searchService.Query(elastic.NewQueryStringQuery(query))
	}

	searchResult, err := searchService.
		//Sort("name", true). // sort by "user" field, ascending
		//From(0).Size(2). // take documents 0-9
		Pretty(true). // pretty print request and response JSON
		Do(r.context) // execute

	if err != nil {
		return nil, err
	}

	for _, item := range searchResult.Each(reflect.TypeOf(domain.Todo{})) {
		todo := item.(domain.Todo)
		output = append(output, todo)
	}

	return output, nil
}
