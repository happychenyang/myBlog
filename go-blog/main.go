package main

import "go-blog/router"

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
