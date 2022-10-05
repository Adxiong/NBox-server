/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-15 00:36:49
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-05 22:46:23
 */
package favorites

import (
	"github.com/gin-gonic/gin"
)

func FavoritesList(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}
