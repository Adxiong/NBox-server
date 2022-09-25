/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-09-25 22:49:26
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-25 22:49:30
 */
package todo

import "github.com/gin-gonic/gin"

func UpdateTodo(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}
