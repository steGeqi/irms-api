package user

import (
	gorms "irms-api.com/project-api/pkg/dao/gorm"
	"time"
)

// 还有几个字段后续再补
type User struct {
	Id         string    `json:"id"`
	GroupId    string    `json:"groupid"`
	InsertDate time.Time `json:"insertdate"`
	UserName   string    `json:"username"`
	RealName   string    `json:"realname"`
	Password   string    `json:"password"`
}

func (u *User) VerifyIdentity(username, password string) bool {
	return false
}

type MemberDao struct {
	conn *gorms.GormConn
}
