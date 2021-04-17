package src

type elasticBody struct {
	Settings settings `json:"settings"`
	Mappings mappings `json:"mappings"`
}

type mappings struct {
	Properties map[string]property `json:"properties"`
}

type property struct {
	Type string `json:"type"`
}

type settings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}
