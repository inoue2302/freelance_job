package components

import (
	"freelance_job/app/models"
	"freelance_job/app/util"
)

type JobComonent struct {
}

const PAGE_NUM = 10

/**
 * GetDetail
 * 指定IDの仕事情報取得
 *
 * @param id int
 * @retrun models.Project|err
 */
func (self *JobComonent) GetDetail(id int) (models.Project, error) {
	//DBコネクション取得
	DbMap := util.GetDb()
	var project models.Project
	//対象データ取得
	err := DbMap.SelectOne(&project, "select * from project where id = ?", id)
	return project, err
}

/**
 * GetMaxCount
 * 仕事情報最大数を取得
 *
 * @return maxCount int
 */
func (self *JobComonent) GetMaxCount() (int64, error) {
	//DBコネクション取得
	DbMap := util.GetDb()
	//最大数取得
	maxCount, err := DbMap.SelectInt("select count(*) from project")
	return maxCount, err
}

/**
 * GetList
 * 仕事情報リストを取得
 *
 * @return []models.Project|error
 */
func (self *JobComonent) GetList(page int) ([]models.Project, error) {
	//開始位置を選出
	offset := page*PAGE_NUM - PAGE_NUM
	//DBコネクション取得
	DbMap := util.GetDb()
	var projects []models.Project
	_, err := DbMap.Select(&projects, "select * from project limit ?, 10", offset)
	return projects, err
}
