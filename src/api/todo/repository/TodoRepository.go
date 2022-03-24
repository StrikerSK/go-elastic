package repository

import (
	"context"
	"encoding/json"
	"fmt"
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

func (r TodoRepositoryStruct) SearchDocument(documentID string) (output entity.Todo, err error) {
	searchResult, err := r.Client.Get().Index(r.IndexName).Id(documentID).Do(r.Context)
	if err != nil {
		log.Printf("Search Index [%s/%s] Error: %v\n", r.IndexName, documentID, err)
		return entity.Todo{}, err
	}

	resolvedStructure, err := searchResult.Source.MarshalJSON()
	if err != nil {
		log.Printf("Search Index [%s/%s] Error: %v\n", r.IndexName, documentID, err)
		return entity.Todo{}, err
	}

	if err = json.Unmarshal(resolvedStructure, &output); err != nil {
		log.Printf("Search in index %s error: %v\n", r.IndexName, err)
		return entity.Todo{}, err
	}

	log.Println("Search successful")
	return
}

func (r TodoRepositoryStruct) GetBySearch(documentID string) (err error) {
	ctx := context.Background()
	termQuery := elastic.NewQueryStringQuery("Create Todo")
	searchResult, err := r.Client.Search().
		Index(r.IndexName). // search in index "twitter"
		Query(termQuery).   // specify the query
		//Sort("name", true). // sort by "user" field, ascending
		//From(0).Size(10). // take documents 0-9
		Pretty(true). // pretty print request and response JSON
		Do(ctx)       // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	for _, item := range searchResult.Each(reflect.TypeOf(entity.Todo{})) {
		t := item.(entity.Todo)
		fmt.Printf("Tweet by %s: %s\n", t.Name, t.Description)

	}

	return
}
