package body

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
