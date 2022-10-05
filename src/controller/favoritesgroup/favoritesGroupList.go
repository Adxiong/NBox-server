/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-31 00:20:25
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-05 22:51:00
 */
package favoritesgroup

import (
	"github.com/gin-gonic/gin"
)

func FavoritesGroupList(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}
