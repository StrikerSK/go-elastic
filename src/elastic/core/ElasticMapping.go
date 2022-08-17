package core

// ElasticMappings - structure mapping all structure field name and type
type ElasticMappings struct {
	Type       string                     `json:"type,omitempty"`
	Properties map[string]ElasticMappings `json:"properties,omitempty"`
}

/*
NewElasticMappings - Constructor to create new ElasticMapping instance.
*/
func NewElasticMappings(propertyType string, propertyMapping map[string]ElasticMappings) *ElasticMappings {
	return &ElasticMappings{
		Type:       propertyType,
		Properties: propertyMapping,
	}
}

func (m *ElasticMappings) addType(key, value string) {
	mapping := ElasticMappings{
		Type: value,
	}

	m.Properties[key] = mapping
	return
}

func (m *ElasticMappings) setMapping(key string, mapping ElasticMappings) {
	m.Properties[key] = mapping
	return
}
