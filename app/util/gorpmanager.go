package util

import (
	"database/sql"
	"freelance_job/app/models"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

var (
	DbMap       *gorp.DbMap
	Transaction *gorp.Transaction
)

/**
 * GetDb
 * DbMap取得
 *
 * @return *gorp.DbMap
 */
func GetDb() *gorp.DbMap {
	DbMap = InitDb()
	return DbMap
}

/**
 * InitDb
 * DbMap生成(シングルトン)
 *
 * @return *gorp.DbMap
 */
func InitDb() *gorp.DbMap {

	if DbMap == nil {
		//コネクション作成
		db, err := sql.Open("mysql", dbInfoString())
		if err != nil {
			panic(err.Error())
		}
		DbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		// 構造体とテーブルの紐付け
		DbMap.AddTableWithName(models.Project{}, "project")
	}
	return DbMap
}

/**
 * dbInfoString
 * DB接続情報取得
 *
 * @return dbInfoText string
 */
func dbInfoString() string {
	dbInfoText, b := revel.Config.String("db.Info")
	if !b {
		panic("database info not found")
	}
	return dbInfoText
}

/**
 * Begin
 * トランザクション開始
 *
 * @return Transaction *gorp.Transaction
 */
func Begin() *gorp.Transaction {
	DbMap := InitDb()
	Transaction, err := DbMap.Begin()
	if err != nil {
		panic(err)
	}
	return Transaction
}

/**
 * Commit
 * トランザクション開始
 *
 * @return void
 */
func Commit() {
	if Transaction == nil {
		return
	}
	err := Transaction.Commit()
	if err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}

/**
 * Rollback
 * ロールバック
 *
 * @return void
 */
func Rollback() {
	if Transaction == nil {
		return
	}
	err := Transaction.Rollback()
	if err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}
