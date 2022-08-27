package exampleDomain

type ExampleStruct struct {
	NestedExampleStruct
}

type NestedExampleStruct struct {
	NestedString string  `json:"nestedString"`
	NestedInt    int64   `json:"nestedNumber"`
	NestedFloat  float32 `json:"nestedFloat"`
}
