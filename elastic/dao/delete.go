package elastic

import (
	"bytes"
	"encoding/json"

	"github.com/gdgtoledo/linneo/elastic"
	"github.com/gdgtoledo/linneo/plants"
	log "github.com/sirupsen/logrus"
)

func (id string) getQuery() map[string]interface{} {
	query := map[string]interface{}{"query": {
		match: {
			id: id,
		},
	}}

	return query
}

// Delete a plant by id
func Delete(id string) (plants.DeleteByIDResult, error) {
	result := plants.DeleteByIDResult{}

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
}
