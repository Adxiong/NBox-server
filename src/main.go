/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 09:58:00
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-07 17:48:12
 */
package main

import (
	"context"
	"nbox/src/bootstrap"
	. "nbox/src/controller"

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

	r.Run(":8000")
}
