/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 23:07:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-07 17:27:57
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
	db.GlobalDb.AutoMigrate(&db.Favorites{})
	db.GlobalDb.AutoMigrate(&db.FavoritesGroup{})
	db.GlobalDb.AutoMigrate(&db.Todo{})
}
