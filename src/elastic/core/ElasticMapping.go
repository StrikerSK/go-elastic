package core

//ElasticMappings - structure mapping all structure field name and type
type ElasticMappings struct {
	Type       string                     `json:"type,omitempty"`
	Properties map[string]ElasticMappings `json:"properties,omitempty"`
}

/*
NewElasticMappings - Constructor to create new ElasticMapping instance.
*/
func NewElasticMappings(propertyType string, propertyMapping map[string]ElasticMappings) *ElasticMappings {
	return &ElasticMappings{
		Type:       propertyType,
		Properties: propertyMapping,
	}
}

func (m *ElasticMappings) addType(key, value string) {
	mapping := ElasticMappings{
		Type: value,
	}

	m.Properties[key] = mapping
	return
}

func (m *ElasticMappings) setMapping(key string, mapping ElasticMappings) {
	m.Properties[key] = mapping
	return
}

/*
CreateElasticObject - Generating of ElasticSearch's simple index model to create. Normally this should work with
nested structs and slices as far as it was tested.
*/
//func CreateElasticObject(customStruct interface{}) *ElasticMappings {
//	structValue := reflect.ValueOf(customStruct)
//
//	outputMapping := NewElasticMappings("", make(map[string]ElasticMappings, structValue.NumField()))
//	for i := 0; i < structValue.NumField(); i++ {
//		fieldObj := structValue.Field(i)
//
//		fieldName := strings.ToLower(structValue.Type().Field(i).Name)
//		fieldKind := fieldObj.Type().Kind()
//
//		fieldType, _ := resolveType(fieldKind.String())
//		nestedMapping := NewElasticMappings(fieldType, make(map[string]ElasticMappings))
//
//		/*
//			In case of 'struct' type, we need to call recursion to resolve nested structure
//		*/
//		if fieldKind == reflect.Struct {
//			nestedMapping.Properties = CreateElasticObject(fieldObj.Interface()).Properties
//		} else if fieldKind == reflect.Slice {
//			/**
//			To resolve slice field we need to find element type of element represented by `fieldObj.Type().Elem()`.
//			Then we need to create new value of this type Calling `reflect.New`. Be aware that this structure will be pointer
//			which need to retrieve the value in this address, done with calling `reflect.Indirect`.
//			**/
//
//			// Elem() - seems to work on slice's elements
// 			fieldElem := fieldObj.Type().Elem()
//			if fieldElem.Kind() == reflect.Struct {
//				sliceStructure := reflect.Indirect(reflect.New(fieldElem)).Interface()
//				nestedMapping.Properties = CreateElasticObject(sliceStructure).Properties
//			} else {
//				tmpType, err := resolveType(fieldElem.Kind().String())
//				if err != nil {
//					log.Println(err)
//				}
//
//				nestedMapping.Type = tmpType
//			}
//		}
//
//		outputMapping.setMapping(fieldName, *nestedMapping)
//	}
//	return outputMapping
//}
//
//func resolveType(input string) (output string, err error) {
//	isInteger := regexp.MustCompile("^u?int\\d{0,2}")
//	isString := regexp.MustCompile("^string$")
//	isFloat := regexp.MustCompile("^float\\d{0,2}$")
//	isBool := regexp.MustCompile("^bool$")
//	isStruct := regexp.MustCompile("^struct$")
//
//	switch {
//	case isString.MatchString(input):
//		output = "text"
//	case isBool.MatchString(input):
//		output = "boolean"
//	case isInteger.MatchString(input):
//		output = "number"
//	case isFloat.MatchString(input):
//		output = "float"
//	case isStruct.MatchString(input):
//		output = "nested"
//	default:
//		err = errors.New("cannot resolve type: " + input)
//	}
//	return
//}
