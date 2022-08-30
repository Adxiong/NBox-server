/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-15 00:37:26
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-31 00:33:41
 */
package todo

import (
	"github.com/gin-gonic/gin"
)

func AddTodo(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}

func UpdateTodo(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}

func DeleteTodo(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}

func TodoList(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}
