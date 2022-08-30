/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-30 23:07:18
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-31 00:30:37
 */
package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
	req := ctx.Request
	fmt.Println(req)
	ctx.JSON(200, "msg:'ok'")
}

func UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "msg:'ok'")

}
func FindUser(ctx *gin.Context) {
	ctx.JSON(200, "msg:'ok'")

}
