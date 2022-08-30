/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 23:05:56
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-31 00:14:59
 */
package bootstrap

import (
	"context"
	. "nbox/src/controller"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context, r *gin.Engine) {
	// 注册路由
	RegisterController(ctx, r)

}
