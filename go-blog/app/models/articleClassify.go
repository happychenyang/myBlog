package models

import (
	orm "go-blog/common/clint"
	"go-blog/common/models"
)

type ArticleClassify struct {
	BaseModel
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	ArticleCount int    `json:"article_count"`
}

func (ArticleClassify) TableName() string {
	return "article_classify"
}

func (u *ArticleClassify) Generate() models.ActiveRecord {
	o := *u
	return &o
}

// GetClassifyByArticle 获取
func (*ArticleClassify) GetClassifyByArticle() (articleClassify []ArticleClassify, err error) {
	var obj = orm.Db.Select("article_classify.*,count(article.id) as article_count").Joins("left join article on article_classify.id = article.classify").Group("article_classify.name").Find(&articleClassify)
	if err = obj.Error; err != nil {
		return
	}
	return
}
