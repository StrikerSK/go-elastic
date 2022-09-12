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
	assert.Equal(t, 0, len(field.Properties))
}

func Test_CreateSliceStringField(t *testing.T) {
	testStruct := struct {
		MultipleStrings []string
	}{}

	elasticStructure := elasticMappingFactory.CreateElasticObject(testStruct)
	field := elasticStructure.Properties["multiplestrings"]
	assert.NotNil(t, field)
	assert.Equal(t, "text", field.Type)
	assert.Equal(t, 0, len(field.Properties))
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
	assert.Equal(t, 0, len(stringProps.Properties))

	numberProps := field.Properties["nestednumber"]
	assert.Equal(t, "number", numberProps.Type)
	assert.Equal(t, 0, len(numberProps.Properties))
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
	assert.Equal(t, 0, len(stringProps.Properties))

	numberProps := field.Properties["nestednumber"]
	assert.Equal(t, "number", numberProps.Type)
	assert.Equal(t, 0, len(numberProps.Properties))
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
	assert.Equal(t, 0, len(stringProps.Properties))

	numberProps := elasticStructure.Properties["nestednumber"]
	assert.Equal(t, "number", numberProps.Type)
	assert.Equal(t, 0, len(numberProps.Properties))
}
