package plants

import dao "github.com/gdgtoledo/linneo/elastic/dao"

// Delete by id
func Delete(id string) (DeleteQueryByIDResult, error) {
	result, err := dao.Delete(id)
	return result, err
}
