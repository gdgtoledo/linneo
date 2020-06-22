package elastic

import (
	"github.com/gdgtoledo/linneo/src/dao"
)

// Search Dao implementation to elastic
func (p Params) Search() (dao.Response, error) {
	var r dao.Response
	var err error
+
	if esClientError != nil {
		return r, err
	}


	return r, err
}
