/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 09:58:00
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-02 00:33:09
 */
package main

import (
	"context"
	"nbox/src/bootstrap"
	. "nbox/src/controller"
	"nbox/src/model/dao/db"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ctx, cancel := context.WithCancel((context.Background()))
	defer cancel()

	bootstrap.MustInit(ctx)
	LoadUserRoutes(r)
	LoadFavoritesRoutes(r)
	LoadTodoRoutes(r)
	LoadToolsRoutes(r)
	user := db.User{
		ID:       1,
		UId:      15727172785,
		Sex:      1,
		Salt:     "asdasdfsad",
		Birthday: time.Now(),
		Name:     "熊友龙",
		Nick:     "adxiong",
		Password: "0417.xyl",
	}
	user.AddUser(ctx)

	user1 := db.User{
		UId:      15727172785,
		Sex:      1,
		Salt:     "asdasdfsad",
		Birthday: time.Now(),
		Name:     "aa",
		Nick:     "adxiong",
		Password: "0417.xyl",
	}
	user1.AddUser(ctx)
	r.Run(":8000")
}
