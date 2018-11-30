package service

import (
	"net/http"
	"strconv"

	"github.com/makki0205/log"

	"github.com/makki0205/config"
)

var Led ledService

type ledService struct {
	url string
}

func init() {
	res := ledService{}
	res.url = config.Env("led_url")
	Led = res
}

func (l *ledService) SetAll(scene uint8) {
	go func() {
		_, err := http.Post(l.url+"/"+strconv.Itoa(int(scene)), "application/json", nil)
		if err != nil {
			log.Err(err)

		}
	}()

}
