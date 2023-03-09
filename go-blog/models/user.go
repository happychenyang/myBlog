package models

import (
	orm "go-blog/database"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

/**
当查到使用  user  会返回列表
没有查到表示空值
*/
func (user *User) FindUserList() (userList []User, err error) {
	if err = orm.Db.Find(&userList).Error; err != nil {
		return
	}
	return
}
