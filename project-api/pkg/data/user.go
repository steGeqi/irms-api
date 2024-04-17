package data

import "time"

// 还有几个字段后续再补
type User struct {
	Id         string    `json:"id"`
	GroupId    string    `json:"groupid"`
	InsertDate time.Time `json:"insertdate"`
	UserName   string    `json:"username"`
	RealName   string    `json:"realname"`
	Password   string    `json:"password"`
}

func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "tb_user"
}
