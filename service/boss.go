package service

import (
	"net/http"

	"github.com/makki0205/config"

	"github.com/makki0205/log"
)

var Boss bossService

type bossService struct {
	url string
}

func init() {
	Boss = bossService{
		url: config.Env("boss_url"),
	}
}
func (b *bossService) send(pass string) {
	go func() {
		_, err := http.Get(b.url + pass)
		if err != nil {
			log.Err(err)
		}
	}()
}

func (b *bossService) Kan() {
	b.send("/boss")
}
