package controllers

import (
	"freelance_job/app/controllers/base"
	"github.com/revel/revel"
)

/**
 * コントローラー開始処理
 */
func init() {
	revel.OnAppStart(base.InitDb)                                         // DBやテーブルの作成
	revel.InterceptMethod((*base.GorpController).Begin, revel.BEFORE)     // transaction開始
	revel.InterceptMethod((*base.GorpController).Commit, revel.AFTER)     // 変更反映
	revel.InterceptMethod((*base.GorpController).Rollback, revel.FINALLY) // 異常時処理
}
