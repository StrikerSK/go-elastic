package exampleService

import (
	exampleDomain "github.com/strikersk/go-elastic/src/api/exampleTodo/domain"
	elasticIndex "github.com/strikersk/go-elastic/src/elastic/core/index"
)

type ExampleTodoService struct {
	indexBuilder elasticIndex.ElasticIndexBuilder
}

func NewExampleTodoService(indexBuilder elasticIndex.ElasticIndexBuilder) ExampleTodoService {
	return ExampleTodoService{
		indexBuilder: indexBuilder,
	}
}

func (ExampleTodoService) GenerateExampleStruct() exampleDomain.ExampleStruct {
	stringField := "String field"
	nestedStruct := exampleDomain.NestedExampleStruct{
		NestedString: "Nested string",
		NestedNumber: 999999999,
		NestedFloat:  0.123456789,
	}

	return exampleDomain.ExampleStruct{
		FieldNumber:         123456789,
		StringField:         stringField,
		NestedStruct:        nestedStruct,
		StringSlice:         []string{"test", "testing", "qa"},
		NestedStructSlice:   []exampleDomain.NestedExampleStruct{nestedStruct, nestedStruct},
		StringPointer:       &stringField,
		StringSlicePointer:  []*string{&stringField, &stringField},
		PointerStructSlice:  []*exampleDomain.NestedExampleStruct{&nestedStruct, &nestedStruct},
		NestedExampleStruct: nestedStruct,
	}
}

func (r ExampleTodoService) CreateExampleIndex() elasticIndex.ElasticIndexBody {
	return r.indexBuilder.BuildIndex(exampleDomain.ExampleStruct{})
}
