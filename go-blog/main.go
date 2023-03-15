package main

import (
	_ "go-blog/app/controller"
	"go-blog/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
