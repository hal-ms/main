package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/main/cnt"
)

func gameRouter(game *gin.RouterGroup) {
	game.GET("/start", cnt.GameStart)
	game.GET("/check", cnt.GameCheck)
	game.GET("/end", cnt.GameEnd)

}