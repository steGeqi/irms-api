package model

// 和用户相关接口的数据模型 及方法

type User struct {
	ID       string `gorms:"column:id"`
	Username string `gorms:"column:username"`
	Password string `gorms:"column:password"`
	Phone    string `gorms:"column:phone"`
}

func (u User) TableName() string {
	//绑定表名
	return "tb_user"
}
