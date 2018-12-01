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
	b.send("wait")
}

func (b *bgmService) Game() {
	b.send(store.Job)
}
func (b *bgmService) Stop() {
	b.send("")
}

func (b *bgmService) IsWare(f bool) {
	if f {
		if store.IsStandby {
			b.Standby()
		} else {
			b.Game()
		}
	} else {
		b.Stop()
	}
}
