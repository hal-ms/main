package main

import (
	"math/rand"
	"time"

	"github.com/makki0205/log"

	"github.com/hal-ms/main/router"
	"github.com/hal-ms/main/service"
	"github.com/hal-ms/main/store"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	store.IsStandby = true
	store.IsWare = false
	service.BGM.Load()
	service.MovingHed.Standby()
	log.ServiceName = "ms-main"
	router.GetRouter().Run(":8000")
}
