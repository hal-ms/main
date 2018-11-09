package cnt

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/main/repo"
)

func GameStart(c *gin.Context) {
	repo.Job.All()

}

func GameCheck(c *gin.Context) {

}

func GameEnd(c *gin.Context) {

}
