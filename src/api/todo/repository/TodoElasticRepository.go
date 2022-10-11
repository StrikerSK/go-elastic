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
		indexName: indexName,
	}
}

func (r TodoElasticRepository) DeleteDocument(ctx context.Context, documentID string) (err error) {
	_, err = r.client.Delete().Index(r.indexName).Id(documentID).Do(ctx)
	return
}

func (r TodoElasticRepository) CreateDocument(ctx context.Context, document todoDomain.Todo) (string, error) {
	return r.insertDocument(ctx, "", document)
}

func (r TodoElasticRepository) UpdateDocument(ctx context.Context, document todoDomain.Todo) error {
	_, err := r.insertDocument(ctx, document.ID, document)
	return err
}

func (r TodoElasticRepository) insertDocument(ctx context.Context, id string, document todoDomain.Todo) (string, error) {
	data, err := json.Marshal(document)
	if err != nil {
		log.Printf("Marshalling document error: %v\n", err)
		return "", err
	}

	response, err := r.client.Index().Index(r.indexName).Id(id).BodyJson(string(data)).Do(ctx)
	if err != nil {
		log.Printf("Insert to index error: %v\n", err)
		return "", err
	}

	log.Println("Document inserted to index")
	return response.Id, nil
}

func (r TodoElasticRepository) ReadDocument(ctx context.Context, documentID string) (output todoDomain.Todo, err error) {
	searchResult, err := r.client.Get().Index(r.indexName).Id(documentID).Do(ctx)
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

func (r TodoElasticRepository) SearchTodos(ctx context.Context, stringQuery string) ([]todoDomain.Todo, error) {
	output := make([]todoDomain.Todo, 0)

	searchService := r.client.Search().Index(r.indexName).Query(elastic.NewQueryStringQuery(stringQuery))
	searchResult, err := searchService.
		Pretty(true).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	for _, item := range searchResult.Each(reflect.TypeOf(todoDomain.Todo{})) {
		todo := item.(todoDomain.Todo)
		output = append(output, todo)
	}

	return output, nil
}
