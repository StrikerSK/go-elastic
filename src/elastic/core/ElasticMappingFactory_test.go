package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var elasticMappingFactory = NewElasticMappingFactory()

type NestedStruct struct {
	NestedString string
	NestedNumber int
}

func Test_CreateStringField(t *testing.T) {
	testStruct := struct {
		SingleString string
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	field := elasticStructure.Properties["singlestring"]
	assert.NotNil(t, field)
	assert.Equal(t, "text", field.Type)
	assert.Empty(t, field.Properties)
}

func Test_CreateSliceStringField(t *testing.T) {
	testStruct := struct {
		MultipleStrings []string
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	field := elasticStructure.Properties["multiplestrings"]
	assert.NotNil(t, field)
	assert.Equal(t, "text", field.Type)
	assert.Empty(t, field.Properties)
}

func Test_CreateStructField(t *testing.T) {
	testStruct := struct {
		Nested NestedStruct
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	field := elasticStructure.Properties["nested"]
	assert.NotNil(t, field)
	assert.Equal(t, "nested", field.Type)
	assert.Equal(t, 2, len(field.Properties))

	stringProps := field.Properties["nestedstring"]
	assert.Equal(t, "text", stringProps.Type)
	assert.Empty(t, stringProps.Properties)

	numberProps := field.Properties["nestednumber"]
	assert.Equal(t, "number", numberProps.Type)
	assert.Empty(t, numberProps.Properties)
}

func Test_CreateStructSliceField(t *testing.T) {
	testStruct := struct {
		Nested []NestedStruct
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	field := elasticStructure.Properties["nested"]
	assert.NotNil(t, field)
	assert.Equal(t, "slice", field.Type)
	assert.Equal(t, 2, len(field.Properties))

	stringProps := field.Properties["nestedstring"]
	assert.Equal(t, "text", stringProps.Type)
	assert.Empty(t, stringProps.Properties)

	numberProps := field.Properties["nestednumber"]
	assert.Equal(t, "number", numberProps.Type)
	assert.Empty(t, numberProps.Properties)
}

func Test_CreateAnonymousField(t *testing.T) {
	testStruct := struct {
		NestedStruct
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	assert.NotNil(t, elasticStructure)
	assert.Equal(t, "", elasticStructure.Type)
	assert.Equal(t, 2, len(elasticStructure.Properties))

	stringProps := elasticStructure.Properties["nestedstring"]
	assert.Equal(t, "text", stringProps.Type)
	assert.Empty(t, stringProps.Properties)

	numberProps := elasticStructure.Properties["nestednumber"]
	assert.Equal(t, "number", numberProps.Type)
	assert.Empty(t, numberProps.Properties)
}

func Test_CreateStructPointerField(t *testing.T) {
	testStruct := struct {
		Nested *NestedStruct
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	stringProps := elasticStructure.Properties["nestedstring"]
	assert.Equal(t, "text", stringProps.Type)
	assert.Empty(t, stringProps.Properties)

	numberProps := elasticStructure.Properties["nestednumber"]
	assert.Equal(t, "number", numberProps.Type)
	assert.Empty(t, numberProps.Properties)
}

func Test_CreateSimplePointerField(t *testing.T) {
	testStruct := struct {
		PointerString *string
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	stringProps := elasticStructure.Properties["pointerstring"]
	assert.Equal(t, "text", stringProps.Type)
	assert.Empty(t, stringProps.Properties)
}

func Test_CreateCombinedFields1(t *testing.T) {
	testStruct := struct {
		SomeString    string
		PointerString *string
		NestedStruct
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	assert.NotNil(t, elasticStructure)
	assert.Equal(t, "", elasticStructure.Type)
	assert.Equal(t, 4, len(elasticStructure.Properties))

	stringProps := elasticStructure.Properties["nestedstring"]
	assert.Equal(t, "text", stringProps.Type)
	assert.Empty(t, stringProps.Properties)

	pointerStringProps := elasticStructure.Properties["pointerstring"]
	assert.Equal(t, "text", pointerStringProps.Type)
	assert.Empty(t, pointerStringProps.Properties)

	numberProps := elasticStructure.Properties["nestednumber"]
	assert.Equal(t, "number", numberProps.Type)
	assert.Empty(t, numberProps.Properties)

	structProps := elasticStructure.Properties["somestring"]
	assert.Equal(t, "text", structProps.Type)
	assert.Empty(t, structProps.Properties)
}
