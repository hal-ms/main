package service

import (
	"bytes"
	"encoding/json"
	"net/http"

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

func (l *ledService) SetAll(r, g, b uint8) {
	req, _ := json.Marshal(struct {
		R uint8 `json:"r"`
		G uint8 `json:"g"`
		B uint8 `json:"b"`
	}{r, g, b})
	_, err := http.Post(l.url+"/", "application/json", bytes.NewReader(req))
	if err != nil {
		log.Err(err)
	}
}
