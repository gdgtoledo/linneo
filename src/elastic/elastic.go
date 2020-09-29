package elastic

import (
	"net/http"

	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/gdgtoledo/linneo/src/dao"
	log "github.com/sirupsen/logrus"
	apmes "go.elastic.co/apm/module/apmelasticsearch"
)

// DataObjectModel type to elastic Dao
type DataObjectModel struct {
	client *es.Client
}

var esClient *es.Client

var esClientError error

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

func get() dao.Interface {
	var edao dao.Interface

	esClient, esClientError = getClient()
	edao = DataObjectModel{client: esClient}

	return edao
}
