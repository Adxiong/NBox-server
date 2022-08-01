/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 23:07:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-01 23:12:09
 */
package bootstrap

import (
	"context"
	"nbox/src/model/dao/db"
)

func MustInit(ctx context.Context) {
	InitMysql()
}

func InitMysql() {
	db.GlobalDb = db.NewDb()
	db.GlobalDb.AutoMigrate(&db.User{})
}
