package main

import "github.com/hal-ms/main/router"

func main() {
	router.GetRouter().Run(":8000")
}
