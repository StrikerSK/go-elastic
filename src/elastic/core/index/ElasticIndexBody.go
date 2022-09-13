package index

import (
	elasticMappings "github.com/strikersk/go-elastic/src/elastic/core/mappings"
	elasticSettings "github.com/strikersk/go-elastic/src/elastic/core/settings"
)

type ElasticIndexBody struct {
	Settings elasticSettings.ElasticSettings `json:"settings"`
	Mappings elasticMappings.ElasticMappings `json:"mappings"`
}

func NewElasticIndexBody() *ElasticIndexBody {
	return &ElasticIndexBody{}
}

func (r *ElasticIndexBody) WithMappings(mappings elasticMappings.ElasticMappings) *ElasticIndexBody {
	r.Mappings = mappings
	return r
}

func (r *ElasticIndexBody) WithSettings(settings elasticSettings.ElasticSettings) *ElasticIndexBody {
	r.Settings = settings
	return r
}

func (r *ElasticIndexBody) WithDefaultSettings() *ElasticIndexBody {
	settings := elasticSettings.NewElasticSettings().WithDefaultShardsCount().WithDefaultReplicasCount()
	return r.WithSettings(*settings)
}
