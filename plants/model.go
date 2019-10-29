package plants

// Model for an item
type Model struct {
	ID      string
	Name    string
	Species string
	Genere  string
	Geo     string
}

// SearchQueryByIndexName struct
type SearchQueryByIndexName struct {
	IndexName string
	Query     map[string]interface{}
}

// SearchQueryByIndexNameResult wraps a search result
type SearchQueryByIndexNameResult map[string]interface{}

// DeleteByIDResult wraps a delete item result
type DeleteQueryByID interface {
	getQuery() map[string]interface{}
}

// DeleteByIDResult wraps a delete item result
type DeleteQueryByIDResult map[string]interface{}
