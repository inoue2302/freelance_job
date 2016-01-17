package base

import (
	"freelance_job/app/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

/**
 * Db初期処理
 */
func InitDb() {
	_ = util.InitDb()
}

/**
 * GorpController
 * Gorpコントローラー
 */
type GorpController struct {
	*revel.Controller
}

/**
 * GorpController
 * トランザクション開始
 */
func (c *GorpController) Begin() revel.Result {
	_ = util.Begin()
	return nil
}

/**
 * GorpController
 * トランザクションコミット
 */
func (c *GorpController) Commit() revel.Result {
	util.Commit()
	return nil
}

/**
 * GorpController
 * ロールバック
 */
func (c *GorpController) Rollback() revel.Result {
	util.Rollback()
	return nil
}
