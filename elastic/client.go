package elastic

import (
	es "github.com/elastic/go-elasticsearch/v8"
	log "github.com/sirupsen/logrus"
)

var esInstance *es.Client

// GetClient returns a client connected to the running elasticseach cluster
func GetClient() (*es.Client, error) {
	if esInstance != nil {
		return esInstance, nil
	}

	cfg := es.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
			"http://elasticsearch2:9200",
			"http://elasticsearch3:9200",
		},
	}
	esClient, err := es.NewClient(cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"config": cfg,
			"error":  err,
		}).Error("Could not obtain an Elasticsearch client")

		return nil, err
	}

	esInstance = esClient

	return esInstance, nil
}
