package models

import (
	"go-blog/app/bean"
	orm "go-blog/common/clint"
	"go-blog/common/models"
)

type Article struct {
	BaseModel
	Name            string          `json:"name"`
	Classify        string          `json:"classify"`
	Introduction    string          `json:"introduction"`
	ArticleImage    string          `json:"article_image"`
	LookNum         int             `json:"look_num"`
	ArticleText     ArticleText     `gorm:"foreignKey:ArticleId"`
	ArticleClassify ArticleClassify `gorm:"foreignKey:Classify"`
}

func (Article) TableName() string {
	return "article"
}

func (u *Article) Generate() models.ActiveRecord {
	o := *u
	return &o
}

// GetArticleList 查询文章列表	没有查到表示空值
func (*Article) GetArticleList(articleWhere *bean.ArticleList) (articleList []Article, total int64, err error) {
	obj := orm.Db
	// 模糊查询
	if articleWhere.Name != "" {
		obj = obj.Where("name like ?", "%"+articleWhere.Name+"%")
	}
	obj = obj.Scopes(Paginate(articleWhere.PageIndex, articleWhere.PageSize)).Order("created_at desc").Find(&articleList)
	if err = obj.Error; err != nil {
		return
	}
	total, err = GetCount(obj)
	return
}

// GetArticleDetail 获取文章详情
func (*Article) GetArticleDetail(articleWhere bean.ArticleDetail) (articleDetail Article, err error) {
	var obj = orm.Db.Preload("ArticleText").Preload("ArticleClassify").Where(articleWhere).First(&articleDetail)
	if err = obj.Error; err != nil {
		return
	}
	return
}
