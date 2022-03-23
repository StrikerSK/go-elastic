package body

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

//Structure mapping all structure field name and type
type ElasticMappings struct {
	Type       string                     `json:"type,omitempty"`
	Properties map[string]ElasticMappings `json:"properties,omitempty"`
}

//NewMappings - Constructor to create new ElasticMapping instance
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

//Generating of ElasticSearches' simple index model to create
func CreateMappingMap(userStruct interface{}) *ElasticMappings {
	structValue := reflect.ValueOf(userStruct)

	outputMapping := NewMappings("", make(map[string]ElasticMappings, structValue.NumField()))
	for i := 0; i < structValue.NumField(); i++ {
		fieldObj := structValue.Field(i)

		fieldName := strings.ToLower(structValue.Type().Field(i).Name)
		fieldType, _ := resolveType(fieldObj.Type().Kind().String())
		resolvedProperty := NewMappings(fieldType, nil)

		//fmt.Printf("Name: %s, Kind: %v\n", fieldName, v.Field(i).Kind().String())
		//In case of 'struct' type, we need to call recursion to resolve nested structure
		fieldKind := fieldObj.Type().Kind().String()
		if fieldKind == "struct" {
			resolvedProperty.Properties = CreateMappingMap(fieldObj.Interface()).Properties
			outputMapping.addType(fieldName, fieldObj.Kind().String())
			outputMapping.Properties[fieldName] = *resolvedProperty
		} else if fieldKind == "slice" {
			t := fieldObj.Kind()
			switch t {
			case reflect.Slice:
				s := reflect.ValueOf(t)

				for i := 0; i < s.Len(); i++ {
					fmt.Println(s.Index(i))
				}
			}
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