/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-15 00:36:49
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-31 00:33:32
 */
package favorites

import (
	"github.com/gin-gonic/gin"
)

func AddFavorites(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}

func UpdateFavorites(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}

func DeleteFavorites(ctx *gin.Context) {

}

func FavoritesList(ctx *gin.Context) {
	ctx.JSON(200, "ok")
}
