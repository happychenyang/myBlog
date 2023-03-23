package bean

//  文章列表
/**
lte：小于等于参数值，validate:“lte=3” (小于等于3)
gte：大于等于参数值，validate:“lte=0,gte=120” (大于等于0小于等于120)
lt：小于参数值，validate:“lt=3” (小于3)
gt：大于参数值，validate:“lt=0,gt=120” (大于0小于120)
len：等于参数值，validate:“len=2”
max：大于等于参数值，validate:“max=2” (大于等于2)
min：小于等于参数值，validate:“min=2,max=10” (大于等于2小于等于10)
ne：不等于，validate:“ne=2” (不等于2)
oneof：只能是列举出的值其中一个，这些值必须是数值或字符串，以空格分隔，如果字符串中有空格，将字符串用单引号包围，validate:“oneof=red green”
*/
type Page struct {
	PageIndex int `form:"PageIndex" json:"PageIndex" binding:"gte=0"  default:"1"`
	PageSize  int `form:"PageSize" json:"PageSize" binding:"gt=0" default:"10"`
}
