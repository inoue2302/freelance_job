package tests

import (
	"encoding/json"
	"freelance_job/app/models"
	simpleJson "github.com/bitly/go-simplejson"
	"github.com/revel/revel/testing"
)

type JobTest struct {
	testing.TestSuite
}

func (t *JobTest) Before() {
	println("Set up")
}

func (t *JobTest) TestJobDetetailNotFound() {
	t.Get("/job/10000000000")
	t.AssertNotFound()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *JobTest) TestJobDetetailBadRequest() {
	t.Get("/job/0")
	t.AssertStatus(400)
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *JobTest) TestJobDetailOk() {
	t.Get("/job/68")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	js, err1 := simpleJson.NewJson(t.ResponseBody)
	t.AssertEqual(nil, err1)

	job, err2 := js.Get("results/Job").Encode()
	t.AssertEqual(nil, err2)

	var project models.Project
	err3 := json.Unmarshal(job, &project)
	t.AssertEqual(nil, err3)
}

func (t *JobTest) TestJobListBadRequest() {
	t.Get("/job/list/0")
	t.AssertStatus(400)
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *JobTest) TestJobList() {
	t.Get("/job/list/1")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	js, err := simpleJson.NewJson(t.ResponseBody)
	t.AssertEqual(nil, err)

	count, _ := js.Get("results/count").Uint64()
	t.Assert(count >= 0)
}

func (t *JobTest) After() {
	println("Tear down")
}
