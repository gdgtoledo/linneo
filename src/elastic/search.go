package elastic

import (
	"github.com/gdgtoledo/linneo/src/dao"
)

// Search to Elastic Dao
func (eDao DataObjectModel) Search(query dao.Query) (dao.Response, error) {
	var r dao.Response
	var err error

	if esClientError != nil {
		return r, err
	}

	return r, err
}
