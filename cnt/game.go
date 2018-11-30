package cnt

import (
	"net/http"
	"time"

	"github.com/hal-ms/main/store"
	"github.com/makki0205/log"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/main/repo"
	"github.com/hal-ms/main/service"
)

func GameStart(c *gin.Context) {
	a := repo.Job.All().IsDone(false).SortByCreated()

	// TODO スタート処理
	// hit画面
	store.Job = a[len(a)-1].Name
	store.IsStandby = false
	service.Hit.Load()
	// ビルのLED
	service.Led.SetAll(14)
	// ムービングヘッド
	service.MovingHed.Player()
	// SE
	c.JSON(http.StatusOK, a[len(a)-1])

}

func GameCheck(c *gin.Context) {

}

func GameEnd(c *gin.Context) {

	// DBの書き換え
	a := repo.Job.All().IsDone(false).SortByCreated()
	if len(a) <= 3 {
		service.CreateRandomJob(1)
	}
	if len(a) <= 0 {
		log.SendSlack("jobがありません")
	}
	now := a[len(a)-1]
	now.Done = true

	// LED 完了アニメーション
	service.Led.SetAll(6)
	// ムービングヘッド
	service.MovingHed.Dool()
	// TODO 終了SE
	service.SE.Clear()
	go func() {
		//演出終了待ち
		time.Sleep(4 * time.Second)
		// LED 完了アニメーション
		service.Led.SetAll(10)
		// ムービングヘッド
		service.MovingHed.Standby()
		//ヒット画面
		store.IsStandby = true
		service.Hit.Load()

		repo.Job.Update(now)
		time.Sleep(400 * time.Microsecond)

		// SE
		service.SE.Kan()
	}()
	c.JSON(http.StatusOK, "ok")
}
