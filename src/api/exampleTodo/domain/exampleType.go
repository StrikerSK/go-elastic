package exampleDomain

type ExampleStruct struct {
	StringField       string                `json:"stringValue"`
	FieldNumber       int                   `json:"numberValue"`
	NestedStruct      NestedExampleStruct   `json:"nestedStruct"`
	NestedStructSlice []NestedExampleStruct `json:"nestedSlice"`
	StringSlice       []string              `json:"stringSlice"`
	StringPointer     *string               `json:"stringPointer"`
	NestedExampleStruct
}

type NestedExampleStruct struct {
	NestedString string  `json:"nestedString"`
	NestedNumber int64   `json:"nestedNumber"`
	NestedFloat  float32 `json:"nestedFloat"`
}
