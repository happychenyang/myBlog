package service

import (
	"go-blog/app/bean"
	model "go-blog/app/models"
)

// GetArticleList 查询基础信息
func GetArticleList(articleWhere *bean.ArticleList) (result []model.Article, total int64, err error) {
	var base model.Article
	result, total, err = base.GetArticleList(articleWhere)
	return
}

// GetArticleDetail 查询文章内容
func GetArticleDetail(articleWhere bean.ArticleDetail) (detail model.Article, err error) {
	var base model.Article
	detail, err = base.GetArticleDetail(articleWhere)
	return
}

// GetArticleClassify 查询文章分类
func GetArticleClassify() (list []model.ArticleClassify, err error) {
	var base model.ArticleClassify
	list, err = base.GetClassifyByArticle()
	return
}
