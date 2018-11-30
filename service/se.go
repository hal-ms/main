package service

import (
	"net/http"

	"github.com/makki0205/log"
)

var SE seService

type seService struct {
}

func (s *seService) send(name string) {
	_, err := http.Get("https://socket.patra.store/emit/se?data=https://hal-iot.net/public/se/" + name + ".mp3")
	if err != nil {
		log.Err(err)
	}
}

func (s *seService) IsWare(f bool) {
	if f {
		s.send("wear")
	} else {
		s.send("notwear")
	}
}
