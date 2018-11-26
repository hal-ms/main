package cnt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsWearing(c *gin.Context) {

	if c.Param("IsWearing") != "true" && c.Param("IsWearing") != "false" {
		c.JSON(http.StatusBadRequest, "パラメータ不正")
		return
	}

	// かぶったor帽子とった処理
	c.JSON(http.StatusOK, "ok")
}
