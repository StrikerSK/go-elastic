package core

type ElasticSettings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}

func NewElasticSettings(shards, replicas int) ElasticSettings {
	return ElasticSettings{
		NumberOfShards:   shards,
		NumberOfReplicas: replicas,
	}
}

func NewDefaultSettings() ElasticSettings {
	return ElasticSettings{
		NumberOfShards:   1,
		NumberOfReplicas: 1,
	}
}
