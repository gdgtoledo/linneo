package elastic

import (
	"net/http"

	es "github.com/elastic/go-elasticsearch/v8"
	log "github.com/sirupsen/logrus"
	apmes "go.elastic.co/apm/module/apmelasticsearch"
)

var esClient *es.Client
var esClientError error

// getElasticsearchClient returns a client connected to the running elasticseach cluster
func getClient() (*es.Client, error) {
	var esInstance *es.Client

	if esInstance != nil {
		return esInstance, nil
	}

	cfg := es.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
			"http://elasticsearch2:9200",
			"http://elasticsearch3:9200",
		},
		Transport: apmes.WrapRoundTripper(http.DefaultTransport),
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

func init() {
	esClient, esClientError = getClient()
}
