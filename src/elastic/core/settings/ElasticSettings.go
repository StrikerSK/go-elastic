package settings

type ElasticSettings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}

func NewElasticSettings() *ElasticSettings {
	return &ElasticSettings{}
}

func (r *ElasticSettings) WithShardsCount(count int) *ElasticSettings {
	r.NumberOfShards = count
	return r
}

func (r *ElasticSettings) WithDefaultShardsCount() *ElasticSettings {
	return r.WithShardsCount(1)
}

func (r *ElasticSettings) WithReplicasCount(count int) *ElasticSettings {
	r.NumberOfReplicas = count
	return r
}

func (r *ElasticSettings) WithDefaultReplicasCount() *ElasticSettings {
	return r.WithReplicasCount(1)
}
