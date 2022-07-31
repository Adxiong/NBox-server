/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 09:58:00
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-07-31 10:43:07
 */
package main

import (
	. "nbox/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	LoadUserRoutes(r)
	LoadFavoritesRoutes(r)
	LoadTodoRoutes(r)
	LoadToolsRoutes(r)
	r.Run(":8000")
}
