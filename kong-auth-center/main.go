package main

import "go-kong-auth-practice/kong-auth-center/router"

func main() {
	router.GetRouter().Run(":8081")
}
