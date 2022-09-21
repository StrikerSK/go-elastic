package mappings

// ElasticMappings - structure mapping all structure field name and type
type ElasticMappings struct {
	Type       string                     `json:"type,omitempty"`
	Properties map[string]ElasticMappings `json:"properties,omitempty"`
}

func NewDefaultMapping() *ElasticMappings {
	return &ElasticMappings{
		Type:       "",
		Properties: make(map[string]ElasticMappings, 0),
	}
}

func (m *ElasticMappings) WithType(typeName string) *ElasticMappings {
	m.Type = typeName
	return m
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

func (m *ElasticMappings) setType(typeValue string) {
	m.Type = typeValue
	return
}

func (m *ElasticMappings) setProperties(mapping *ElasticMappings) {
	tmpMap := make(map[string]ElasticMappings, 0)
	for key, element := range m.Properties {
		tmpMap[key] = element
	}

	for key, element := range mapping.Properties {
		tmpMap[key] = element
	}

	m.Properties = tmpMap
	return
}

func (m *ElasticMappings) changeProperties(key string, mapping ElasticMappings) {
	m.Properties[key] = mapping
	return
}
