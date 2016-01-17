package controllers

import (
	"freelance_job/app/components"
	"freelance_job/app/controllers/base"
	"freelance_job/app/models"
	"github.com/revel/revel"
)

/**
 * JobController
 * ジョブコントローラー
 */
type JobController struct {
	base.ApiBaseController
}

/**
 * JobList
 * 案件リストレスポンス構造体
 */
type JobList struct {
	MaxCount int64            `json:"max_count"`
	Jobs     []models.Project `json:"jobs"`
}

/**
 * JobDetail
 * 案件詳細情報レスポンス構造体
 */
type JobDetail struct {
	Job models.Project `josn:job`
}

/**
 * Detail
 * 案件詳細情報取得アクション
 *
 * @param id int
 * @return json
 */
func (c *JobController) Detail(id int) revel.Result {
	c.Validation.Required(id)
	if c.Validation.HasErrors() {
		return c.HandleBadRequestError("request param error")
	}
	jobComponent := new(components.JobComonent)
	project, err := jobComponent.GetDetail(id)
	if err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	res := base.Response{JobDetail{project}}
	return c.RenderJson(res)
}

/**
 * List
 * 案件リスト取得アクション
 *
 * @param page int
 * @return json
 */
func (c *JobController) List(page int) revel.Result {
	c.Validation.Required(page)
	if c.Validation.HasErrors() {
		return c.HandleBadRequestError("request param error")
	}
	jobComponent := new(components.JobComonent)
	maxCount, err_count := jobComponent.GetMaxCount()

	if err_count != nil {
		c.HandleBadRequestError(err_count.Error())
	}

	projects, err_list := jobComponent.GetList(page)
	if err_list != nil {
		c.HandleInternalServerError(err_list.Error())
	}
	res := base.Response{JobList{maxCount, projects}}
	return c.RenderJson(res)
}
