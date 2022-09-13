package core

import (
	"errors"
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
	structType := reflect.TypeOf(inputStruct)

	outputMapping := NewDefaultMapping()
	for i := 0; i < structValue.NumField(); i++ {
		fieldObject := structValue.Field(i)
		fieldObjectKind := fieldObject.Type().Kind()
		fieldObjectName := strings.ToLower(structValue.Type().Field(i).Name)

		isFieldAnonymous := structType.Field(i).Anonymous

		fieldType, err := r.resolveType(fieldObjectKind.String())
		if err != nil {
			log.Println(err)
		}

		nestedMapping := NewDefaultMapping().WithType(fieldType)

		/*
			In case of 'struct' type, we need to call recursion to resolve nested structure
		*/

		switch fieldObjectKind {
		case reflect.Struct:
			structInterface := fieldObject.Interface()
			if isFieldAnonymous {
				r.mapStructType(outputMapping, structInterface)
			} else {
				r.mapStructType(nestedMapping, structInterface)
			}
		case reflect.Slice:
			/**
			To resolve slice field we need to find element type of element represented by `fieldObj.Type().Elem()`.
			Then we need to create new value of this type Calling `reflect.New`. Be aware that this structure will be pointer
			which need to retrieve the value in this address, done with calling `reflect.Indirect`.
			**/

			// Elem() - seems to work on slice's elements
			sliceElem := fieldObject.Type().Elem()
			if sliceElem.Kind() == reflect.Struct {
				sliceStruct := r.mapPointerType(fieldObject)
				r.mapStructType(nestedMapping, sliceStruct)
			} else {
				r.mapStandardType(nestedMapping, sliceElem.Kind())
			}
		case reflect.Pointer:
			pointerObject := r.mapPointerType(fieldObject)
			pointerObjectKind := reflect.TypeOf(pointerObject).Kind()
			if pointerObjectKind == reflect.Struct {
				r.mapStructType(outputMapping, pointerObject)
			} else {
				r.mapStandardType(nestedMapping, pointerObjectKind)
			}
		}

		if !isFieldAnonymous {
			outputMapping.changeProperties(fieldObjectName, *nestedMapping)
		}
	}
	return outputMapping
}

// mapStandardType - Add mapping of standard type of checked field
func (r ElasticMappingFactory) mapStandardType(mapper *ElasticMappings, kind reflect.Kind) {
	sliceType, err := r.resolveType(kind.String())
	if err != nil {
		log.Println(err)
	}

	mapper.setType(sliceType)
}

// mapStructType - Add mapping of struct type of checked field
func (r ElasticMappingFactory) mapStructType(mapper *ElasticMappings, inputStruct interface{}) {
	pointerProperties := r.CreateElasticObject(inputStruct)
	mapper.setProperties(pointerProperties)
}

// mapPointerType - Resolves pointer field type of struct to regular struct
func (r ElasticMappingFactory) mapPointerType(inputStruct reflect.Value) interface{} {
	pointerElem := inputStruct.Type().Elem()
	pointerValue := reflect.New(pointerElem)
	return reflect.Indirect(pointerValue).Interface()
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
