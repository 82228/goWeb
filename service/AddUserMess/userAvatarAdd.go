package AddUserMess

import "webb/dao"

func AddAvatar(account, avatarUrl string) error {
	return dao.AddUserAvatar(account, avatarUrl)
}
