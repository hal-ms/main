package main

import (
	"github.com/hal-ms/main/router"
	"github.com/hal-ms/main/service"
	"github.com/makki0205/log"
)

func main() {
	service.BGM.Standby()
	log.ServiceName = "ms-main"
	router.GetRouter().Run(":8000")
}
