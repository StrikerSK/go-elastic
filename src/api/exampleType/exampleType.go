package exampleType

type exampleStruct struct {
	FirstString  string                `json:"firstString"`
	FirstNumber  string                `json:"firstNumber"`
	NestedStruct nestedExampleStruct   `json:"nestedStruct"`
	NestedSlice  []nestedExampleStruct `json:"somethingElse"`
	SliceString  []string              `json:"sliceString"`
}

type nestedExampleStruct struct {
	NestedString string  `json:"nestedString"`
	NestedInt    int64   `json:"nestedNumber"`
	NestedFloat  float32 `json:"nestedFloat"`
}
