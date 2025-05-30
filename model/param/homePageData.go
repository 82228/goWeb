package param

import "webb/model"

type HomePageData struct {
	CurUser    model.User   //当前账号
	UserList   []model.User //所有账号信息
	CurserRole int
}
