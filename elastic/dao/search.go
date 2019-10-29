package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gdgtoledo/linneo/elastic"
	"github.com/gdgtoledo/linneo/plants"
	log "github.com/sirupsen/logrus"
)

// Search executes a query in the proper index
func Search(query plants.SearchQueryByIndexName) (plants.SearchQueryByIndexNameResult, error) {
	result := plants.SearchQueryByIndexNameResult{}

	esClient, err := elastic.GetClient()

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
		esClient.Search.WithIndex(query.IndexName),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
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
