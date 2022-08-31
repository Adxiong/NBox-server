/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 23:05:56
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-31 23:23:47
 */
package bootstrap

import (
	"context"
	"log"
	. "nbox/src/controller"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context) {
	MustInit(ctx)

	r := gin.New()

	// 注册路由
	RegisterController(ctx, r)

	if err := r.Run(":8000"); err != nil {
		log.Printf("startup service failed, err:%v\n", err)
	}
}
