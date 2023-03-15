package models

import (
	orm "go-blog/common/clint"
	"go-blog/common/models"
)

type Menu struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	IconName string `json:"icon_name"`
	Url      string `json:"url"`
}

func (Menu) TableName() string {
	return "menu"
}

func (u *Menu) Generate() models.ActiveRecord {
	o := *u
	return &o
}

// FindMenuList 查询菜单列表	没有查到表示空值
func (*Rotation) FindMenuList() (menuList []Menu, err error) {
	var obj = orm.Db.Find(&menuList)
	if err = obj.Error; err != nil {
		return
	}
	return
}
