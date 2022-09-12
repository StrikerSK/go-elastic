package exampleDomain

type ExampleStruct struct {
	FirstString  string                `json:"stringValue"`
	FirstNumber  string                `json:"numberValue"`
	NestedStruct NestedExampleStruct   `json:"nestedStruct"`
	NestedSlice  []NestedExampleStruct `json:"nestedSlice"`
	SliceString  []string              `json:"stringSlice"`
	NestedExampleStruct
}

type NestedExampleStruct struct {
	NestedString string  `json:"nestedString"`
	NestedInt    int64   `json:"nestedNumber"`
	NestedFloat  float32 `json:"nestedFloat"`
}
