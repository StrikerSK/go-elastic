package core

// ElasticMappings - structure mapping all structure field name and type
type ElasticMappings struct {
	Type       string                     `json:"type,omitempty"`
	Properties map[string]ElasticMappings `json:"properties,omitempty"`
}

func NewDefaultMapping(mapSize int) *ElasticMappings {
	if mapSize == 0 {
		mapSize = 1
	}

	return &ElasticMappings{
		Type:       "",
		Properties: make(map[string]ElasticMappings, mapSize),
	}
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

func (m *ElasticMappings) setProperties(properties map[string]ElasticMappings) {
	m.Properties = properties
	return
}

func (m *ElasticMappings) setPropertiesFromMapping(mapping *ElasticMappings) {
	tmpMap := make(map[string]ElasticMappings)
	for key, element := range m.Properties {
		tmpMap[key] = element
	}

	for key, element := range mapping.Properties {
		tmpMap[key] = element
	}

	m.Properties = tmpMap
	return
}

func (m *ElasticMappings) addType(key, value string) {
	mapping := ElasticMappings{
		Type: value,
	}

	m.Properties[key] = mapping
	return
}

func (m *ElasticMappings) changeProperties(key string, mapping ElasticMappings) {
	m.Properties[key] = mapping
	return
}
