package cnt

import (
	"net/http"
	"time"

	"github.com/hal-ms/main/store"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/log"
	"github.com/hal-ms/main/repo"
	"github.com/hal-ms/main/service"
)

func GameStart(c *gin.Context) {
	a := repo.Job.All().IsDone(false).SortByCreated()

	// TODO スタート処理
	store.Job = a[len(a)-1].Name
	store.IsStandby = false

	service.Hit.Load()
	service.Led.SetAll(255, 0, 0)
	c.JSON(http.StatusOK, a[len(a)-1])

}

func GameCheck(c *gin.Context) {

}

func GameEnd(c *gin.Context) {
	service.Led.SetAll(255, 255, 0)
	a := repo.Job.All().IsDone(false).SortByCreated()
	if len(a) <= 3 {
		service.CreateRandomJob(2)
	}
	if len(a) <= 0 {
		log.SendSlack("jobがありません")
		panic("jobがありません")
	}
	now := a[len(a)-1]
	now.Done = true
	// TODO end処理

	repo.Job.Update(now)
	go func() {
		store.IsStandby = true
		time.Sleep(4000 + time.Second)
		service.Led.SetAll(0, 0, 255)
	}()
	c.JSON(http.StatusOK, "ok")
}
