/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 10:21:23
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-07-31 10:43:25
 */
package controller

import "github.com/gin-gonic/gin"

// LoadUserRoutes 加载user组路由
func LoadUserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, "msg:'ok'")
		})
	}

}

// LoadFavoritesRoutes 加载favorites组路由
func LoadFavoritesRoutes(r *gin.Engine) {
	favorites := r.Group("/favorites")
	{
		favorites.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, "msg:'ok'")
		})
	}
}

// LoadTodoRoutes 加载Todo组路由
func LoadTodoRoutes(r *gin.Engine) {
	todo := r.Group("/todo")
	{
		todo.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, "msg:'ok'")
		})
	}
}

// LoadtoolsRoutes 加载tools组路由
func LoadToolsRoutes(r *gin.Engine) {
	tools := r.Group("/tools")
	{
		tools.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, "msg:'ok'")
		})
	}
}
