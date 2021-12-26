package main

import "go-kong-auth-practice/kong-sys-service/router"

func main() {
	router.GetRouter().Run(":8087")
}
