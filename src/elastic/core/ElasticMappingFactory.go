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

/*
CreateElasticObject - Generating of ElasticSearch's simple index model to create. Normally this should work with
nested structs and slices as far as it was tested.
*/
func (r ElasticMappingFactory) CreateElasticObject(inputStruct interface{}) *ElasticMappings {
	structValue := reflect.ValueOf(inputStruct)
	structTypeOf := reflect.TypeOf(inputStruct)

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
			structProperties := r.CreateElasticObject(fieldObj.Interface())
			if isFieldAnonymous {
				outputMapping.setPropertiesFromMapping(structProperties)
			} else {
				nestedMapping.setPropertiesFromMapping(structProperties)
			}
		} else if fieldKind == reflect.Slice {
			/**
			To resolve slice field we need to find element type of element represented by `fieldObj.Type().Elem()`.
			Then we need to create new value of this type Calling `reflect.New`. Be aware that this structure will be pointer
			which need to retrieve the value in this address, done with calling `reflect.Indirect`.
			**/

			// Elem() - seems to work on slice's elements
			sliceElem := fieldObj.Type().Elem()
			if sliceElem.Kind() == reflect.Struct {
				sliceStruct := reflect.Indirect(reflect.New(sliceElem)).Interface()
				sliceProperties := r.CreateElasticObject(sliceStruct)
				nestedMapping.setPropertiesFromMapping(sliceProperties)
			} else {
				sliceType, err := r.resolveType(sliceElem.Kind().String())
				if err != nil {
					log.Println(err)
				}

				nestedMapping.setType(sliceType)
			}
		} else if fieldKind == reflect.Pointer {
			pointerValue := reflect.New(fieldObj.Type().Elem())
			pointerObject := reflect.Indirect(pointerValue).Interface()

			if reflect.TypeOf(pointerObject).Kind() == reflect.Struct {
				pointerProperties := r.CreateElasticObject(pointerObject)
				outputMapping.setPropertiesFromMapping(pointerProperties)
			} else {
				kindValue := reflect.TypeOf(pointerObject).Kind().String()
				sliceType, err := r.resolveType(kindValue)
				if err != nil {
					log.Println(err)
				}

				nestedMapping.setType(sliceType)
			}

		}

		if !isFieldAnonymous {
			outputMapping.changeProperties(fieldName, *nestedMapping)
		}
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
