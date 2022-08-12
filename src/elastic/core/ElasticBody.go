package core

type ElasticBody struct {
	Settings ElasticSettings `json:"settings"`
	Mappings ElasticMappings `json:"mappings"`
}

func NewDefaultElasticBody(mappings ElasticMappings) ElasticBody {
	return ElasticBody{
		Settings: NewDefaultSettings(),
		Mappings: mappings,
	}
}

func NewElasticBody(settings ElasticSettings, mappings ElasticMappings) ElasticBody {
	return ElasticBody{
		Settings: settings,
		Mappings: mappings,
	}
}

type ElasticIndexBuilder struct {
	mappingFactory ElasticMappingFactory
}

func NewElasticIndexBuilder(mappingFactory ElasticMappingFactory) ElasticIndexBuilder {
	return ElasticIndexBuilder{
		mappingFactory: mappingFactory,
	}
}

func (r ElasticIndexBuilder) BuildIndex(structure interface{}) ElasticBody {
	mapping := *r.mappingFactory.CreateElasticObject(structure)
	return NewDefaultElasticBody(mapping)
}
