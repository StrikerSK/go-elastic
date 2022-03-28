package elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/strikersk/go-elastic/src/api/todo/entity"
	"github.com/strikersk/go-elastic/src/api/todo/repository"
	"log"
	"os"
	"time"
)

func GetElasticInstance() {
	log.Println("ElasticSearch initialization")

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(os.Getenv("ELASTIC_URL")),
		elastic.SetHealthcheckInterval(5*time.Second),
	)

	if err != nil {
		log.Printf("ElasticSearch initialization error: %s\n", err)
		os.Exit(1)
	}

	elasticConfig := ElasticConfiguration{
		ElasticClient: client,
		Context:       context.Background(),
	}

	elasticConfig.InitializeIndex(entity.TodoIndex, entity.CreateTodoIndexBody())

	repository.SetTodoRepository(repository.NewElasticRepository(client, elasticConfig.Context))

	log.Println("ElasticSearch initialization completed!")
}
