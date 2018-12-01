package cnt

import (
	"net/http"
	"strconv"
	"time"

	"github.com/hal-ms/main/cnt/util"

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
	// BGM
	service.BGM.Game()
	// ビルのLED
	service.Led.SetAll(21)
	// ムービングヘッド
	service.MovingHed.Game(1)
	// SE
	service.SE.Start()

	c.JSON(http.StatusOK, a[len(a)-1])

}

func GameCheck(c *gin.Context) {
	scene, err := strconv.Atoi(c.Param("scene"))
	if err != nil {
		util.BadRequest(err.Error(), c)
	}
	switch scene {
	case 1:
		// ビル
		service.Led.SetAll(20)
		// スポットライト
		service.MovingHed.Game(2)
		break
	case 2:
		// ビル
		service.Led.SetAll(19)
		// スポットライト
		service.MovingHed.Game(3)
		break
	}
	go func() {
		time.Sleep(800 * time.Microsecond)
		service.SE.Game(scene)
	}()
}

func GameEnd(c *gin.Context) {

	// DBの書き換え
	a := repo.Job.All().IsDone(false).SortByCreated()
	if len(a) <= 1 {
		service.CreateRandomJob(1)
	}
	if len(a) <= 0 {
		log.SendSlack("jobがありません")
	}
	now := a[len(a)-1]
	now.Done = true
	service.BGM.Stop()
	// LED 完了アニメーション
	service.Led.SetAll(0)
	// ムービングヘッド
	service.MovingHed.Dool()

	go func() {
		time.Sleep(3 * time.Second)
		// BGM停止
		service.BGM.Stop()
		// 終了SE
		service.SE.Clear()
		//演出終了待ち
		time.Sleep(4 * time.Second)
		// LED 完了アニメーション
		service.Led.SetAll(16)
		// ムービングヘッド
		service.MovingHed.Standby()
		//ヒット画面
		store.IsStandby = true
		service.Hit.Load()

		repo.Job.Update(now)
		time.Sleep(400 * time.Microsecond)

		// SE
		service.SE.Kan()
		time.Sleep(1 * time.Second)
		service.BGM.Standby()
	}()
	c.JSON(http.StatusOK, "ok")
}
