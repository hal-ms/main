package service

import (
	"net/http"
	"strconv"

	"github.com/hal-ms/main/store"

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

func (s *seService) Kan() {
	s.send("caaaan")
}
func (s *seService) Clear() {
	s.send("clear")
}

func (s *seService) Game(scene int) {
	s.send("cut_" + store.Job + "_" + strconv.Itoa(scene))
}
func (s *seService) IsWare(f bool) {
	if f {
		s.send("wear")
	} else {
		s.send("notwear")
	}
}
