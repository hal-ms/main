package main

import (
	"github.com/hal-ms/main/router"
	"github.com/makki0205/log"
)

func main() {
	log.ServiceName = "ms-main"
	router.GetRouter().Run(":8000")
}
