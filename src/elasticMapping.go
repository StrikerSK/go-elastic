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

var mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"records":{
			"properties":{
				"name":{
					"type":"keyword"
				},
				"age":{
					"type":"keyword"
				},
				"email":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"updated":{
					"type":"date"
				}
			}
		}
	}
}`
