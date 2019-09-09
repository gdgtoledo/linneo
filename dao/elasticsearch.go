package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	es "github.com/elastic/go-elasticsearch/v8"
	log "github.com/sirupsen/logrus"
	apmes "go.elastic.co/apm/module/apmelasticsearch"
)

var esInstance *es.Client

// SearchResult wraps a search result
type SearchResult map[string]interface{}

// getElasticsearchClient returns a client connected to the running elasticseach cluster
func getElasticsearchClient() (*es.Client, error) {
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

// Search executes a query in the proper index
func Search(ctx context.Context, indexName string, query map[string]interface{}) (SearchResult, error) {
	result := SearchResult{}

	esClient, err := getElasticsearchClient()
	if err != nil {
		return result, err
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error encoding Elasticsearch query")

		return result, err
	}

	log.WithFields(log.Fields{
		"query": fmt.Sprintf("%s", query),
	}).Debug("Elasticsearch query")

	res, err := esClient.Search(
		esClient.Search.WithIndex(indexName),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
		esClient.Search.WithContext(ctx),
	)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error getting response from Elasticsearch")

		return result, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Error parsing error response body from Elasticsearch")

			return result, err
		}

		log.WithFields(log.Fields{
			"status": res.Status(),
			"type":   e["error"].(map[string]interface{})["type"],
			"reason": e["error"].(map[string]interface{})["reason"],
		}).Error("Error getting response from Elasticsearch")

		return result, fmt.Errorf(
			"Error getting response from Elasticsearch. Status: %s, Type: %s, Reason: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"])
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error parsing response body from Elasticsearch")

		return result, err
	}

	log.WithFields(log.Fields{
		"status": res.Status(),
		"hits":   int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		"took":   int(result["took"].(float64)),
	}).Debug("Response information")

	return result, nil
}
