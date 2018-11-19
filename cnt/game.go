package cnt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/log"
	"github.com/hal-ms/main/repo"
	"github.com/hal-ms/main/service"
)

func GameStart(c *gin.Context) {
	repo.Job.All()

}

func GameCheck(c *gin.Context) {

}

func GameEnd(c *gin.Context) {
	a := repo.Job.All().IsDone(false).SortByCreated()
	if len(a) <= 0 {
		log.SendSlack("jobがありません")
		panic("jobがありません")
	}
	if len(a) == 3 {
		service.CreateRandomJob(2)
	}
	now := a[len(a)-1]
	now.Done = true
	// TODO end処理

	repo.Job.Update(now)
	c.JSON(http.StatusOK, gin.H{"job": a[len(a)-2].Name})
}
