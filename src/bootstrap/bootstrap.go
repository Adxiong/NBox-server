/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 23:07:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-24 22:10:56
 */
package bootstrap

import (
	"context"
	"io"
	"log"
	"nbox/src/model/dao/db"
	"os"
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

func InitLog() {
	f, err := os.OpenFile("log/nbox.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()

	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
