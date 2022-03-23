package elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/strikersk/go-elastic/src/types"
	"log"
)

type ElasticConfiguration struct {
	ElasticClient *elastic.Client
	Context       context.Context
}

func (ec ElasticConfiguration) InitializeIndex(indexName string, indexBody []byte) {
	exists, err := ec.indexExists(indexName)
	if err != nil {
		return
	}

	if exists {
		log.Printf("Index [%s] Initialize: index already exists\n", indexName)
		if err = ec.deleteIndex(indexName); err != nil {
			return
		}
	}

	if err = ec.createIndex(indexName, indexBody); err != nil {
		return
	}

	log.Printf("Index [%s] Recreated\n", indexName)
}

func (ec ElasticConfiguration) createIndex(indexName string, indexBody []byte) error {
	if _, err := ec.ElasticClient.CreateIndex(indexName).Body(string(indexBody)).Do(ec.Context); err != nil {
		log.Printf("Index [%s] Create: %v\n", indexName, err)
		return err
	}

	log.Printf("Index [%s] Create: success\n", indexName)
	return nil
}

func (ec ElasticConfiguration) indexExists(indexName string) (bool, error) {
	exists, err := ec.ElasticClient.IndexExists(indexName).Do(ec.Context)
	if err != nil {
		log.Printf("Index [%s] existance check: %v\n", indexName, err)
		return exists, err
	} else {
		return exists, nil
	}
}

func (ec ElasticConfiguration) deleteIndex(indexName string) error {
	if _, err := ec.ElasticClient.DeleteIndex(indexName).Do(ec.Context); err != nil {
		log.Printf("Index [%s] Delete: %v\n", indexName, err)
		return err
	}

	return nil
}

func (ec ElasticConfiguration) SearchDocument(documentID string, targetClass types.MarshallingInterface) error {
	indexName := targetClass.GetIndexName()
	searchResult, err := ec.ElasticClient.
		Get().
		Index(indexName).
		Id(documentID).
		Do(ec.Context)

	//TODO Create solution to transfer status code
	//This might be always not found
	if err != nil {
		log.Printf("Search Index [%s/%s] Error: %v\n", indexName, documentID, err)
		return err
	}

	resolvedStructure, err := searchResult.Source.MarshalJSON()
	if err != nil {
		log.Printf("Search Index [%s/%s] Error: %v\n", indexName, documentID, err)
		return err
	}

	if err = targetClass.UnmarshalItem(resolvedStructure); err != nil {
		log.Printf("Search Index [%s/%s] Error: %v\n", indexName, documentID, err)
		return err
	}

	log.Printf("Search Index [%s/%s]: success\n", indexName, documentID)
	return nil
}

func (ec ElasticConfiguration) InsertDocument(documentID string, input types.MarshallingInterface) (string, error) {
	indexName := input.GetIndexName()

	dataJSON, err := input.MarshalItem()
	if err != nil {
		log.Printf("Insert Document to Index [%s]: %v\n", indexName, err)
		return "", err
	}

	contentBody := string(dataJSON)
	replyCustom, err := ec.ElasticClient.Index().
		Index(indexName).
		Id(documentID).
		BodyJson(contentBody).
		Do(ec.Context)

	if err != nil {
		log.Printf("Insert Document to Index [%s]: %s\n", indexName, err)
		return "", err
	}

	log.Printf("Insert Document to Index [%s]: success\n", indexName)
	return replyCustom.Id, nil
}

func (ec ElasticConfiguration) DeleteDocument(documentID, indexName string) {
	_, err := ec.ElasticClient.
		Delete().
		Index(indexName).
		Id(documentID).
		Do(ec.Context)

	if err != nil {
		log.Printf("Index Document [%s/%s] Delete: %v", indexName, documentID, err)
		return
	}
}
