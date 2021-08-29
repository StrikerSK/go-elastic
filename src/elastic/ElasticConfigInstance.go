package elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"sync"
	"time"
)

var elasticLock = &sync.Mutex{}
var elasticConfiguration *ElasticConfiguration

func GetElasticInstance() *ElasticConfiguration {
	//To prevent expensive lock operations
	//This means that the cacheConnection field is already populated
	if elasticConfiguration == nil {
		elasticLock.Lock()
		defer elasticLock.Unlock()

		//Only one goroutine can create the singleton instance.
		if elasticConfiguration == nil {
			var configuration ElasticConfiguration
			log.Println("Creating ElasticSearch instance")
			client, err := elastic.NewClient(
				elastic.SetSniff(false),
				elastic.SetURL(os.Getenv("ELASTIC_URL")),
				elastic.SetHealthcheckInterval(5*time.Second),
			)

			if err != nil {
				log.Printf("ElasticSearch Initialization error: %s\n", err)
				os.Exit(1)
			}

			configuration.ElasticClient = client
			configuration.Context = context.Background()
			elasticConfiguration = &configuration
			log.Println("ElasticSearch initialized...")
		} else {
			log.Println("ElasticSearch instance already created!")
		}
	} else {
		//log.Println("Application Cache instance already created!")
	}

	return elasticConfiguration
}
