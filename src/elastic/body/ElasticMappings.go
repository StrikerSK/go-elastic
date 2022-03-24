package body

import (
	"errors"
	"reflect"
	"regexp"
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
		fieldKind := fieldObj.Type().Kind()

		fieldType, _ := resolveType(fieldKind.String())
		resolvedProperty := NewMappings(fieldType, make(map[string]ElasticMappings))

		//fmt.Printf("Name: %s, Kind: %v\n", fieldName, v.Field(i).Kind().String())
		//In case of 'struct' type, we need to call recursion to resolve nested structure
		if fieldKind == reflect.Struct {
			nestedStructure := fieldObj.Interface()
			resolvedProperty.Properties = CreateMappingMap(nestedStructure).Properties
			outputMapping.addType(fieldName, fieldObj.Kind().String())
			outputMapping.Properties[fieldName] = *resolvedProperty
		} else if fieldKind == reflect.Slice {
			/**
			To resolve slice field we need to find element type of element represented by `fieldObj.Type().Elem()`.
			Then we need to create new value of this type Calling `reflect.New`. Be aware that this structure will be pointer
			which need to retrieve the value in this address, done with calling `reflect.Indirect`.
			**/
			fieldElem := fieldObj.Type().Elem()
			if fieldElem.Kind() == reflect.Struct {
				sliceStructure := reflect.Indirect(reflect.New(fieldElem)).Interface()
				resolvedProperty.Properties = CreateMappingMap(sliceStructure).Properties
				outputMapping.addType(fieldName, fieldObj.Kind().String())
				outputMapping.Properties[fieldName] = *resolvedProperty
			} else {
				tmpType, _ := resolveType(fieldElem.Kind().String())
				resolvedProperty.Type = tmpType
			}
		}

		outputMapping.Properties[fieldName] = *resolvedProperty
	}
	return outputMapping
}

func resolveType(input string) (output string, err error) {
	isInteger := regexp.MustCompile("^[u]?int\\d{0,2}")
	isString := regexp.MustCompile("^string$")
	isFloat := regexp.MustCompile("^float\\d{0,2}$")
	isBool := regexp.MustCompile("^bool$")
	isStruct := regexp.MustCompile("^struct$")

	switch {
	case isString.MatchString(input):
		output = "text"
	case isBool.MatchString(input):
		output = "boolean"
	case isInteger.MatchString(input):
		output = "number"
	case isFloat.MatchString(input):
		output = "float"
	case isStruct.MatchString(input):
		output = "nested"
	default:
		err = errors.New("cannot resolve type: " + input)
	}
	return
}
