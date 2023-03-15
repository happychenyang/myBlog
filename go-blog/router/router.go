package router

import (
	"github.com/gin-gonic/gin"
	"go-blog/middleware"
	"reflect"
	"strings"
)

// Route 路由参数
type Route struct {
	Middleware gin.HandlerFunc //中间件
	path       string          //请求路径
	httpMethod string          //方法 get  post
	Method     reflect.Value   //方法路由
	Args       []reflect.Type  //参数方法
}

// Routes 路由组
var Routes []Route

func InitRouter() *gin.Engine {
	//自动路由
	r := gin.Default()
	Bind(r)
	return r
}

/**
Register 注册控制器
方法如果是需要get访问   在方法前缀加上 Get_xxx 表示Get访问 否则默认使用POST
*/
func Register(controller interface{}, middleware gin.HandlerFunc) bool {
	ctrName := reflect.TypeOf(controller)
	//fmt.Println("ctrName=", ctrName)
	module := ctrName.String()
	if strings.Contains(module, ".") {
		module = module[strings.Index(module, ".")+1:]
	}
	v := reflect.ValueOf(controller)
	//fmt.Println("module=", module)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name
		path := "/" + module + "/" + action
		params := make([]reflect.Type, 0, v.NumMethod())
		httpMethod := "POST"
		//判断字符串是否包含 _get 如果是表示 get请求
		//HasPrefix 判断字符串 s 是否以 prefix 开头
		if strings.HasPrefix(action, "Get") {
			httpMethod = "GET"
		}

		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
			//fmt.Println("params-name=", method.Type().In(j))
		}
		//中间键配对
		route := Route{path: path, Method: method, Args: params, httpMethod: httpMethod, Middleware: middleware}
		Routes = append(Routes, route)
	}
	return true
}

// Bind 绑定基本路由
func Bind(e *gin.Engine) {
	e.Use(middleware.Cors()) //cors 中间键
	for _, route := range Routes {
		if route.httpMethod == "GET" {
			if route.Middleware != nil {
				e.GET(route.path, route.Middleware, match(route.path, route))
			} else {
				e.GET(route.path, match(route.path, route))
			}
		}
		if route.httpMethod == "POST" {
			if route.Middleware != nil {
				e.POST(route.path, route.Middleware, match(route.path, route))
			} else {
				e.POST(route.path, match(route.path, route))
			}
		}
	}
}

// match 这个是根据path匹配对应的方法
func match(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		if len(fields) < 3 {
			return
		}
		if len(fields) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // gin.Context
			route.Method.Call(arguments)
		}
	}
}
