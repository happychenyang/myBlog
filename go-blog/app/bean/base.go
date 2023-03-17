package bean

//  文章列表
type Page struct {
	PageIndex int `form:"page_index" json:"page_index" binding:"gt=0"  default:"1"`
	PageSize  int `form:"page_size" json:"page_size" binding:"gt=0" default:"10"`
}
