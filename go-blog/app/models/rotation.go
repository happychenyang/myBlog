package models

import (
	orm "go-blog/common/clint"
	"go-blog/common/models"
)

type Rotation struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	BlogImage string `json:"blog_image"`
}

func (Rotation) TableName() string {
	return "rotation"
}

func (u *Rotation) Generate() models.ActiveRecord {
	o := *u
	return &o
}

/**
没有查到表示空值
*/
func (*Rotation) FindRotationList() (rotationList []Rotation, err error) {
	var obj = orm.Db.Find(&rotationList)
	if err = obj.Error; err != nil {
		return
	}
	return
}
