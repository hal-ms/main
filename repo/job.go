package repo

import (
	"net/http"
	"strings"

	"github.com/makki0205/log"

	"github.com/gin-gonic/gin/json"

	"github.com/gin-gonic/gin"

	"github.com/hal-ms/job/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Job = jobRepo{db.C("job")}

type jobRepo struct {
	C *mgo.Collection
}

func (j *jobRepo) Store(job model.Job) model.Job {
	job.ID = bson.NewObjectId()
	j.C.Insert(&job)
	return job
}

func (j *jobRepo) Update(job model.Job) {
	j.C.UpdateId(job.ID, job)
}

func (j *jobRepo) FindByID(id string) *model.Job {
	if !bson.IsObjectIdHex(id) {
		return nil
	}
	job := model.Job{}
	err := j.C.FindId(bson.ObjectIdHex(id)).One(&job)
	if err != nil {
		return nil
	}
	return &job
}
func (j *jobRepo) All() model.Jobs {
	var jobs = make(model.Jobs, 0)
	j.C.Find(nil).All(&jobs)
	return jobs
}

func (j *jobRepo) Create(jobName, UserName string) {
	b, err := json.Marshal(gin.H{"name": jobName, "user_name": UserName})
	if err != nil {
		log.Err(err)
	}
	http.Post("https://hal-iot.net/api/job/masui", "application/json", strings.NewReader(string(b)))
}
