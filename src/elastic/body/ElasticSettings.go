package body

type ElasticSettings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}

func NewDefaultSettings() ElasticSettings {
	return ElasticSettings{
		NumberOfShards:   1,
		NumberOfReplicas: 1,
	}
}
