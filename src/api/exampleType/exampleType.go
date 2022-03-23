package exampleType

type exampleStruct struct {
	FirstString  string                `json:"firstString"`
	FirstNumber  string                `json:"firstNumber"`
	NestedStruct nestedExampleStruct   `json:"nestedStruct"`
	NestedSlice  []nestedExampleStruct `json:"somethingElse"`
}

type nestedExampleStruct struct {
	NestedString string `json:"nestedString"`
	NestedNumber int64  `json:"nestedNumber"`
}
