package bean

// ArticleDetail 获取详情参数
type ArticleDetail struct {
	Id int `form:"articleId" json:"articleId" binding:"gt=0"`
}

// ArticleList 文章列表
type ArticleList struct {
	Page
	Name string `form:"name" json:"name"`
}
