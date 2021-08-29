package elastic

import (
	"errors"
	"reflect"
	"strings"
)

type ElasticBody struct {
	Settings ElasticSettings `json:"settings"`
	Mappings ElasticMappings `json:"mappings"`
}

func NewElasticBody(settings ElasticSettings, sm ElasticMappings) ElasticBody {
	return ElasticBody{
		Settings: settings,
		Mappings: sm,
	}
}

//Structure mapping all structure field name and type
type ElasticMappings struct {
	Type       string                     `json:"type,omitempty"`
	Properties map[string]ElasticMappings `json:"properties,omitempty"`
}

//Constructor to create new ElasticMappings
func NewMappings(propType string, propertiesMapping map[string]ElasticMappings) *ElasticMappings {
	return &ElasticMappings{
		Type:       propType,
		Properties: propertiesMapping,
	}
}

func (m *ElasticMappings) addType(key, value string) {
	m.Properties[key] = ElasticMappings{Type: value}
	return
}

type ElasticSettings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}

func NewDefaultSettings() ElasticSettings {
	return ElasticSettings{
		NumberOfShards:   1,
		NumberOfReplicas: 1,
	}
}

//Generating of ElasticSearches' simple index model to create
func CreateMappingMap(userStruct interface{}) *ElasticMappings {
	v := reflect.ValueOf(userStruct)
	typeOfS := v.Type()

	outputMapping := NewMappings("", make(map[string]ElasticMappings, v.NumField()))
	for i := 0; i < v.NumField(); i++ {
		fieldName := strings.ToLower(typeOfS.Field(i).Name)
		fieldType, _ := resolveType(v.Field(i).Type().Kind().String())
		resolvedProperty := NewMappings(fieldType, nil)

		//fmt.Printf("Name: %s, Kind: %v\n", fieldName, v.Field(i).Kind().String())
		//In case of 'struct' type, we need to call recursion to resolve nested structure
		if v.Field(i).Type().Kind().String() == "struct" {
			nestedMapping := CreateMappingMap(v.Field(i).Interface())
			resolvedProperty.Properties = nestedMapping.Properties
			outputMapping.addType(fieldName, v.Field(i).Kind().String())
			outputMapping.Properties[fieldName] = *resolvedProperty
		}

		outputMapping.Properties[fieldName] = *resolvedProperty
	}
	return outputMapping
}

func resolveType(input string) (output string, err error) {
	switch input {
	case "string":
		output = "text"
	case "bool":
		output = "boolean"
	case "uint8":
		output = "number"
	case "uint16":
		output = "number"
	case "uint32":
		output = "number"
	case "uint64":
		output = "number"
	case "int8":
		output = "number"
	case "int16":
		output = "number"
	case "int32":
		output = "number"
	case "int64":
		output = "number"
	case "struct":
		output = "nested"
	default:
		err = errors.New("cannot resolve type " + input)
	}
	return
}
