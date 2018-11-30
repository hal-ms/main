package service

import (
	"encoding/json"
	"net/http"

	"github.com/hal-ms/main/store"

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
	"priest":     {"https://hal-iot.net/public/hit/no_wear/Monk.jpg", "https://hal-iot.net/public/hit/wear/Monk.jpg"},
	"programmer": {"https://hal-iot.net/public/hit/no_wear/Programmer.jpg", "https://hal-iot.net/public/hit/wear/Programmer.jpg"},
}

func (i *hitService) Load() {
	if store.IsStandby {
		i.send("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTZlmqxCcUMq5X3xmUl_y7dYJDBtOLUhG0io8oHwh4-5j6IXdr4ag")
	}
	i.Set(store.Job, store.IsWare)
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
	http.Get("https://socket.patra.store/emit/hit?data=" + string(b))
}
