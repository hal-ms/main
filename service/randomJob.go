package service

import (
	"math/rand"
	"time"

	"github.com/hal-ms/main/repo"
)

func CreateRandomJob(c int) {
	go func() {
		for i := 0; i < c; i++ {
			time.Sleep(3 * time.Second)
			repo.Job.Create(repo.JobCategory.All()[rand.Int()%len(repo.JobCategory.All())].Name, userNames[rand.Int()%len(userNames)])
		}
	}()

}

var userNames = []string{
	"たま",
	"たかやま",
	"はらがみ",
	"ますい",
	"みうら",
	"すずき",
	"しいたけ",
	"しめじ",
	"にんじん",
	"たまねぎ",
	"織田信長",
	"豊臣秀吉",
	"徳川家康",
	"神様",
	"体臭ヒーロー",
	"レモン",
	"あっぷる",
	"自宅警備員",
	"しゃちく",
	"ぽち",
	"やまだ",
	"さとうとしお",
	"ジョン・スミス",
	"Michel",
	"きつね",
	"インドネシア",
	"平等院鳳凰堂",
}
