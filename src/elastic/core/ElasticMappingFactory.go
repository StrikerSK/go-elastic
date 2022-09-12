package core

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

type ElasticMappingFactory struct{}

func NewElasticMappingFactory() ElasticMappingFactory {
	return ElasticMappingFactory{}
}

//func (r ElasticMappingFactory) createMap(customStruct interface{}) map[string]interface{} {
//	return map[string]interface{}{
//		"properties": r.CreateElasticObject(customStruct),
//	}}

/*
CreateElasticObject - Generating of ElasticSearch's simple index model to create. Normally this should work with
nested structs and slices as far as it was tested.
*/
func (r ElasticMappingFactory) CreateElasticObject(customStruct interface{}) *ElasticMappings {
	structValue := reflect.ValueOf(customStruct)
	structTypeOf := reflect.TypeOf(customStruct)

	outputMapping := NewDefaultMapping(structValue.NumField())
	for i := 0; i < structValue.NumField(); i++ {
		fieldObj := structValue.Field(i)
		fieldKind := fieldObj.Type().Kind()
		fieldName := strings.ToLower(structValue.Type().Field(i).Name)

		isFieldAnonymous := structTypeOf.Field(i).Anonymous

		fieldType, err := r.resolveType(fieldKind.String())
		if err != nil {
			fmt.Println(err)
		}

		nestedMapping := NewDefaultMapping(0)
		nestedMapping.setType(fieldType)

		/*
			In case of 'struct' type, we need to call recursion to resolve nested structure
		*/
		if fieldKind == reflect.Struct {
			properties := r.CreateElasticObject(fieldObj.Interface())
			if !isFieldAnonymous {
				nestedMapping.setPropertiesFromMapping(properties)
			} else {
				outputMapping.setPropertiesFromMapping(properties)
				return outputMapping
			}
		} else if fieldKind == reflect.Slice {
			/**
			To resolve slice field we need to find element type of element represented by `fieldObj.Type().Elem()`.
			Then we need to create new value of this type Calling `reflect.New`. Be aware that this structure will be pointer
			which need to retrieve the value in this address, done with calling `reflect.Indirect`.
			**/

			// Elem() - seems to work on slice's elements
			fieldElem := fieldObj.Type().Elem()
			if fieldElem.Kind() == reflect.Struct {
				sliceStructure := reflect.Indirect(reflect.New(fieldElem)).Interface()
				properties := r.CreateElasticObject(sliceStructure)
				nestedMapping.setPropertiesFromMapping(properties)
			} else {
				tmpType, err := r.resolveType(fieldElem.Kind().String())
				if err != nil {
					log.Println(err)
				}

				nestedMapping.setType(tmpType)
			}
		}
		outputMapping.changeProperties(fieldName, *nestedMapping)
	}
	return outputMapping
}

func (ElasticMappingFactory) resolveType(input string) (output string, err error) {
	isInteger := regexp.MustCompile("^u?int\\d{0,2}")
	isString := regexp.MustCompile("^string$")
	isFloat := regexp.MustCompile("^float\\d{0,2}$")
	isBool := regexp.MustCompile("^bool$")
	isStruct := regexp.MustCompile("^struct$")
	isSlice := regexp.MustCompile("^slice$")

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
	case isSlice.MatchString(input):
		output = "slice"
	default:
		err = errors.New("cannot resolve type: " + input)
	}
	return
}
