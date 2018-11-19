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
}
