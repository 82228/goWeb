package DelUserMess

import "webb/dao"

func Delete(id int) error {
	return dao.DeleteById(id)
}
