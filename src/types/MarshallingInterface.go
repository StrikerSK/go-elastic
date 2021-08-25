package types

//Important during the call to the ElasticSearchServer
type MarshallingInterface interface {
	MarshalItem() ([]byte, error)
	UnmarshalItem([]byte) error
}
