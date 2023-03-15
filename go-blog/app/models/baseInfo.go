package models

import (
	orm "go-blog/common/clint"
	"go-blog/common/models"
)

type BaseInfo struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	BlogImage    string `json:"blog_image"`
	Message      string `json:"message"`
	MessageImage string `json:"message_image"`
}

func (BaseInfo) TableName() string {
	return "base_info"
}

func (u *BaseInfo) Generate() models.ActiveRecord {
	o := *u
	return &o
}

/**
没有查到表示空值
*/
func (*BaseInfo) FindBaseInfo() (userList BaseInfo, err error) {
	var userObj = orm.Db.First(&userList)
	if err = userObj.Error; err != nil {
		return
	}
	return
}

/**
CheckAccount 校验账后是否存在
*/
//func (*BaseInfo) CheckAccount(userEntity BaseInfo) bool {
//	var row = orm.Db.Where(&userEntity).First(&userEntity)
//	if row.Error != nil {
//		return false
//	}
//	return true
//}
