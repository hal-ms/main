package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/makki0205/log"
)

var Hit hitService

type hitService struct {
}

func init() {
	Hit = hitService{}
}

var hitData = map[string][]string{
	"cook":       {"https://hal-iot.net/public/hit/no_wear/Cook.jpg", "https://hal-iot.net/public/hit/wear/Cook.jpg"},
	"pianist":    {"https://hal-iot.net/public/hit/no_wear/Pianist.jpg", "https://hal-iot.net/public/hit/wear/Pianist.jpg"},
	"carpenter":  {"https://hal-iot.net/public/hit/no_wear/Carpenter.jpg", "https://hal-iot.net/public/hit/wear/Carpenter.jpg"},
	"monk":       {"https://hal-iot.net/public/hit/no_wear/Monk.jpg", "https://hal-iot.net/public/hit/wear/Monk.jpg"},
	"programmer": {"https://hal-iot.net/public/hit/no_wear/Programmer.jpg", "https://hal-iot.net/public/hit/wear/Programmer.jpg"},
}

func (i *hitService) Set(job string, isWare bool) {
	if isWare {
		i.send(hitData[job][1])
	} else {
		i.send(hitData[job][0])
	}
}
func (i *hitService) send(image string) {
	req := struct {
		Img string `json:"img"`
	}{
		Img: image,
	}
	b, err := json.Marshal(req)
	if err != nil {
		log.Err(err)
	}
	fmt.Println("https://socket.patra.store/emit/hit?data=" + string(b))
	http.Get("https://socket.patra.store/emit/hit?data=" + string(b))
}
