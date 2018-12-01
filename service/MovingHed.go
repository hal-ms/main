package service

import (
	"net/http"
	"strconv"

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
	fstr := "false"
	if f {
		fstr = "true"
	}
	m.send(m.url + "/wear/" + fstr)
}
func (m *movingHedService) Dool() {
	m.send(m.url + "/clear")
}
func (m *movingHedService) Player() {
	m.send(m.url + "/player")

}

func (m *movingHedService) Standby() {
	m.send(m.url + "/standby")
}
func (m *movingHedService) Game(scene int) {
	m.send(m.url + "/game/" + strconv.Itoa(scene))
}

func (m *movingHedService) send(url string) {
	go func() {
		_, err := http.Get(url)
		if err != nil {
			log.Err(err)
		}
	}()

}
