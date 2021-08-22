package src

import (
	"errors"
	"reflect"
)

type elasticBody struct {
	Settings settings `json:"settings"`
	Mappings mappings `json:"mappings"`
}

type mappings struct {
	Properties map[string]property `json:"properties"`
}

func NewMappings(initSize int) *mappings {
	initializedMap := make(map[string]property, initSize)
	return &mappings{Properties: initializedMap}
}

func (m *mappings) addType(key, value string) {
	m.Properties[key] = property{Type: value}
	return
}

type property struct {
	Type string `json:"type"`
}

type settings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}

func CreateMapping(s interface{}) *mappings {
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	mappingMap := NewMappings(v.NumField())
	for i := 0; i < v.NumField(); i++ {
		resolvedType, _ := resolveType(v.Field(i).Type().Kind().String())
		mappingMap.addType(typeOfS.Field(i).Name, resolvedType)
	}

	return mappingMap
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
