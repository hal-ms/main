package service

import (
	"net/http"

	"github.com/makki0205/config"
	"github.com/makki0205/log"
)

var MovingHed movingHedService

func init() {
	res := movingHedService{}
	res.url = config.Env("moving_url")
	MovingHed = res
}

type movingHedService struct {
	url string
}

func (m *movingHedService) IsWear(f bool) {

}

func (m *movingHedService) Player(f bool) {

}

func (m *movingHedService) Standby(f bool) {

}
func (m *movingHedService) Game(f bool) {

}

func (m *movingHedService) Dool(f bool) {

}
func (m *movingHedService) send(url string) {
	_, err := http.Get(url)
	if err != nil {
		log.Err(err)
	}
}
