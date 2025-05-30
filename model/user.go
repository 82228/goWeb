package model

type User struct {
	Id               int `json:"id"`
	Account          string
	Password         string
	Confrom_Password string
	Role             int //1为管理员，0为普通用户
	Gender           string
	Birthday         string
	Province         string
	Address          string
	AvatarURL        string
	UserName         string
	City             string
}
