package plants

import (
	dao "github.com/gdgtoledo/linneo/elastic/dao"
)

// Search by query with an indexName
func Search(searchQueryByIndexName SearchQueryByIndexName) (SearchQueryByIndexNameResult, error) {
	response, err := dao.Search(searchQueryByIndexName)
	return response, err
}
