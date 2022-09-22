package exampleDomain

type ExampleStruct struct {
	FieldNumber        int                    `json:"numberValue"`
	StringField        string                 `json:"stringValue"`
	NestedStruct       NestedExampleStruct    `json:"nestedStruct"`
	StringSlice        []string               `json:"stringSlice"`
	NestedStructSlice  []NestedExampleStruct  `json:"nestedSlice"`
	StringPointer      *string                `json:"stringPointer"`
	StringSlicePointer []*string              `json:"stringSlicePointer"`
	PointerStructSlice []*NestedExampleStruct `json:"pointerStructSlice"`
	NestedExampleStruct
}

type NestedExampleStruct struct {
	NestedString string  `json:"nestedString"`
	NestedNumber int64   `json:"nestedNumber"`
	NestedFloat  float32 `json:"nestedFloat"`
}
