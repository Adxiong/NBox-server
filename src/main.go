/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 09:58:00
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-04 23:36:50
 */
package main

import (
	"context"
	"fmt"
	"nbox/src/bootstrap"
	. "nbox/src/controller"
	utils "nbox/src/library"

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

	msg := utils.GenLogMessage(utils.LOG_ERROR, "Response_Err")
	fmt.Println(msg)
	r.Run(":8000")
}
