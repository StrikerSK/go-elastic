package repository

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/strikersk/go-elastic/src/api/todo/entity"
	"log"
	"reflect"
)

const TodoIndex = "todos"

type TodoRepositoryStruct struct {
	Client    *elastic.Client
	Context   context.Context
	IndexName string
}

func NewTodoRepository(client *elastic.Client, context context.Context) *TodoRepositoryStruct {
	return &TodoRepositoryStruct{
		Client:    client,
		Context:   context,
		IndexName: TodoIndex,
	}
}

func (r TodoRepositoryStruct) DeleteDocument(documentID string) (err error) {
	_, err = r.Client.Delete().Index(r.IndexName).Id(documentID).Do(r.Context)
	return
}

func (r TodoRepositoryStruct) InsertDocument(documentID string, document entity.Todo) (string, error) {
	dataJSON, err := json.Marshal(document)
	if err != nil {
		log.Printf("Marshalling document error: %v\n", err)
		return "", err
	}

	response, err := r.Client.Index().Index(r.IndexName).Id(documentID).BodyJson(string(dataJSON)).Do(r.Context)
	if err != nil {
		log.Printf("Insert to index error: %v\n", err)
		return "", err
	}

	log.Println("Document inserted to index")
	return response.Id, nil
}

func (r TodoRepositoryStruct) SearchDocumentByID(documentID string) (output entity.Todo, err error) {
	searchResult, err := r.Client.Get().Index(r.IndexName).Id(documentID).Do(r.Context)
	if err != nil {
		log.Printf("Searching error: %v\n", err)
		return entity.Todo{}, err
	}

	resolvedStructure, err := searchResult.Source.MarshalJSON()
	if err != nil {
		log.Printf("Marshalling error: %v\n", err)
		return entity.Todo{}, err
	}

	if err = json.Unmarshal(resolvedStructure, &output); err != nil {
		log.Printf("Unmarshalling error: %v\n", err)
		return entity.Todo{}, err
	}

	log.Println("Search successful")
	return
}

func (r TodoRepositoryStruct) GetByStringQuery(stringQuery []string) ([]entity.Todo, error) {
	output := make([]entity.Todo, 0)

	searchService := r.Client.Search().
		Index(r.IndexName) // search in index "todos"

	for _, query := range stringQuery {
		searchService = searchService.Query(elastic.NewQueryStringQuery(query))
	}

	searchResult, err := searchService.
		//Sort("name", true). // sort by "user" field, ascending
		//From(0).Size(2). // take documents 0-9
		Pretty(true). // pretty print request and response JSON
		Do(r.Context) // execute

	if err != nil {
		return nil, err
	}

	for _, item := range searchResult.Each(reflect.TypeOf(entity.Todo{})) {
		todo := item.(entity.Todo)
		output = append(output, todo)
	}

	return output, nil
}
