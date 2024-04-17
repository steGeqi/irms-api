package dao

import (
	"fmt"
	//"fmt"
	"irms-api.com/project-api/pkg/dao/gorms"
	"irms-api.com/project-api/pkg/model"
)

//type UserDao interface {
//	CreateUser(ctx context.Context, user *model.User) error
//}

type UserDao struct {
	//conn *gorms.GormConn
}

//	func NewUserDao() *UserDao {
//		return &UserDao{
//			conn: gorms.New(),
//		}
//	}
func (u UserDao) VerifyUser(username, password string) bool {
	var userList model.User
	dao.DB.Where("username = ?", username).First(&userList)
	fmt.Println(userList.Password)
	if userList.Password == password {
		return true
	} else {
		return false
	}
	//return false
}
