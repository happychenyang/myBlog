package service

import (
	model "go-blog/app/models"
)

//GetBaseInfo  查询基础信息
func GetBaseInfo() (result model.BaseInfo, err error) {
	var base model.BaseInfo
	result, err = base.FindBaseInfo()
	return
}

// GetRotationList 获取轮播列表
func GetRotationList() (result []model.Rotation, err error) {
	var base model.Rotation
	result, err = base.FindRotationList()
	return
}

// GetMenuList 获取菜单列表
func GetMenuList() (result []model.Menu, err error) {
	var base model.Rotation
	result, err = base.FindMenuList()
	return
}
