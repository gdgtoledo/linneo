package service

import (
	"github.com/gdgtoledo/linneo/src/dao"
	"github.com/gdgtoledo/linneo/src/elastic"
)

// Search in Data Access Object
func Search(query string) (dao.Response, error) {
	var d dao.Dao
	d = elastic.Params{Query: query}
	r, err := d.Search()
	return r, err
}

// Delete by ID in Data Access Object
func Delete(id string) (dao.Response, error) {
	var d dao.Dao
	d = elastic.Params{ID: id}
	r, err := d.Delete()
	return r, err
}
