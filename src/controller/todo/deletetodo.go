/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-15 00:37:26
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-25 22:49:53
 */
package todo

import (
	"github.com/gin-gonic/gin"
)

func DeleteTodo(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}
