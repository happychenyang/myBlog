package models

import (
	orm "go-blog/common/clint"
	"go-blog/common/models"
)

type Verify struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	AppKey string `json:"app_key"`
}

func (Verify) TableName() string {
	return "verify"
}

func (u *Verify) Generate() models.ActiveRecord {
	o := *u
	return &o
}

/**
CheckAccount 校验账后是否存在
*/
func (*Verify) CheckAccount(userEntity Verify) bool {
	var row = orm.Db.Where(&userEntity).First(&userEntity)
	if row.Error != nil {
		return false
	}
	return true
}
