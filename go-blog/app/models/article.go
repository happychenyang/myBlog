package models

import (
	orm "go-blog/common/clint"
	"go-blog/common/models"
)

type Article struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	ArticleImage string `json:"article_image"`
	LookNum      int    `json:"look_num"`
	CreatedAt    string `json:"created_at"`
}

func (Article) TableName() string {
	return "article"
}

func (u *Article) Generate() models.ActiveRecord {
	o := *u
	return &o
}

// GetArticleList 查询文章列表	没有查到表示空值
func (*Article) GetArticleList() (articleList []Article, err error) {
	var obj = orm.Db.Order("created_at desc").Find(&articleList)
	if err = obj.Error; err != nil {
		return
	}
	return
}
