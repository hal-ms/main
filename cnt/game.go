package cnt

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/log"
	"github.com/hal-ms/main/repo"
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
	now := a[len(a)-1]
	now.Done = true
	// TODO end処理

	repo.Job.Update(now)
}
