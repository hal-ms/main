package cnt

import (
	"fmt"
	"net/http"

	"github.com/makki0205/log"

	"github.com/hal-ms/main/service"

	"github.com/hal-ms/main/store"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/config"
)

func IsWearing(c *gin.Context) {
	var flg = false
	if c.Param("IsWearing") != "true" && c.Param("IsWearing") != "false" {
		c.JSON(http.StatusBadRequest, "パラメータ不正")
		return
	}
	if c.Param("IsWearing") == "true" {
		flg = true
	}

	// かぶったor帽子とった処理
	store.IsWare = flg
	service.Hit.Load()
	_, err := http.Post(config.Env("game_url")+"/is_wearing/"+c.Param("IsWearing"), "", nil)
	if err != nil {
		log.Err(err)
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, "ok")
}
