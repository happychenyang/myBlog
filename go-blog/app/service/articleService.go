package service

import model "go-blog/app/models"

//GetBaseInfo  查询基础信息
func GetArticleList() (result []model.Article, err error) {
	var base model.Article
	result, err = base.GetArticleList()
	return
}
