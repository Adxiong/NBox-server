/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 09:58:00
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-31 00:03:25
 */
package main

import (
	"context"
	"nbox/src/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ctx, cancel := context.WithCancel((context.Background()))
	defer cancel()

	bootstrap.MustInit(ctx)
	bootstrap.Start(ctx, r)
	r.Run(":8000")
}
