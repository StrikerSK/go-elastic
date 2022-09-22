package index

import (
	"encoding/json"
	elasticMappings "github.com/strikersk/go-elastic/src/elastic/core/mappings"
)

type ElasticIndexBuilder struct {
	mappingFactory elasticMappings.ElasticMappingFactory
}

func NewElasticIndexBuilder(mappingFactory elasticMappings.ElasticMappingFactory) ElasticIndexBuilder {
	return ElasticIndexBuilder{
		mappingFactory: mappingFactory,
	}
}

func (r ElasticIndexBuilder) BuildIndex(structure interface{}) ElasticIndexBody {
	mapping := *r.mappingFactory.CreateElasticObject(structure)
	return *NewElasticIndexBody().WithDefaultSettings().WithMappings(mapping)
}

func (r ElasticIndexBuilder) BuildAndMarshallIndex(structure interface{}) ([]byte, error) {
	indexData, err := json.Marshal(r.BuildIndex(structure))
	if err != nil {
		return nil, err
	}

	return indexData, nil
}
