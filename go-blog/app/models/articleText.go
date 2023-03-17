package models

import (
	"go-blog/common/models"
)

type ArticleText struct {
	Id             int32  `json:"id"`
	ArticleId      string `json:"article_id"`
	ArticleContent string `json:"article_content"`
}

func (ArticleText) TableName() string {
	return "article_text"
}

func (u *ArticleText) Generate() models.ActiveRecord {
	o := *u
	return &o
}
