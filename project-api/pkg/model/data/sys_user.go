package data

// 和用户相关接口的数据模型 及方法

type User struct {
	ID       int64
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Phone    string `gorm:"column:phone"`
}

func (u User) TableName() string {
	//绑定表名
	return "sys_user"
}
