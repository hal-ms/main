package service

import (
	"fmt"
	"net/http"

	"github.com/hal-ms/main/store"
	"github.com/makki0205/log"
)

var BGM bgmService

type bgmService struct {
}

func (s *bgmService) send(name string) {
	fmt.Println("https://socket.patra.store/emit/bgm?data=https://hal-iot.net/public/bgm/" + name + ".mp3")
	_, err := http.Get("https://socket.patra.store/emit/bgm?data=https://hal-iot.net/public/bgm/" + name + ".mp3")
	if err != nil {
		log.Err(err)
	}
}

func (b *bgmService) Standby() {
	if store.IsWare {
		b.send("wait")
	} else {
		b.send("nowait")
	}
}
func (b *bgmService) StandbyNoWare() {
}

func (b *bgmService) Game() {
	if store.IsWare {
		b.send(store.Job)

	} else {
		b.Stop()
	}
}
func (b *bgmService) Stop() {
	b.send("")
}

func (b *bgmService) Load() {
	if store.IsStandby {
		b.Standby()
	} else {
		b.Game()
	}
}
