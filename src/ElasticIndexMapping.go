package src

import (
	"errors"
	"reflect"
	"strings"
)

type elasticBody struct {
	Settings settings          `json:"settings"`
	Mappings structureMappings `json:"mappings"`
}

//Structure mapping all structure field name and type
type structureMappings struct {
	Type       string                       `json:"type,omitempty"`
	Properties map[string]structureMappings `json:"properties,omitempty"`
}

//Constructor to create new structureMappings
func NewMappings(propType string, propertiesMapping map[string]structureMappings) *structureMappings {
	return &structureMappings{
		Type:       propType,
		Properties: propertiesMapping,
	}
}

func (m *structureMappings) addType(key, value string) {
	m.Properties[key] = structureMappings{Type: value}
	return
}

type settings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}

//Generating of ElasticSearches' simple index model to create
func CreateMappingMap(userStruct interface{}) *structureMappings {
	v := reflect.ValueOf(userStruct)
	typeOfS := v.Type()

	outputMapping := NewMappings("", make(map[string]structureMappings, v.NumField()))
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
