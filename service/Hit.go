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

func (i *hitService) Set() {

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
