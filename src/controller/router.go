/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 10:21:23
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-31 00:35:09
 */
package controller

import (
	"context"
	"nbox/src/controller/favorites"
	"nbox/src/controller/favoritesgroup"
	"nbox/src/controller/todo"
	"nbox/src/controller/user"
	"nbox/src/library/request"

	"github.com/gin-gonic/gin"
)

func RegisterController(ctx context.Context, r *gin.Engine) {
	for key, item := range routes {
		r.Handle(item.Method, key, item.Handle...)
	}
}

type RouterHandler struct {
	Method string
	Handle []gin.HandlerFunc
	Filter func()
}

type HandleFuncs []gin.HandlerFunc

var routes = map[string]RouterHandler{
	"/user/add": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			user.AddUser,
		},
	},
	"/user/update": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			user.UpdateUser,
		},
	},
	"/user/user": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			user.FindUser,
		},
	},

	"/favoritesgroup/add": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			favoritesgroup.AddFavoritesGroup,
		},
	},
	"/favoritesgroup/update": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			favoritesgroup.UpdateFavoritesGroup,
		},
	},
	"/favoritesgroup/delete": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			favoritesgroup.DeleteFavoritesGroup,
		},
	},
	"/favoritesgroup/list": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			favoritesgroup.FavoritesGroupList,
		},
	},

	"/favorites/add": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			favorites.AddFavorites,
		},
	},
	"/favorites/update": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			favorites.UpdateFavorites,
		},
	},
	"/favorites/delete": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			favorites.DeleteFavorites,
		},
	},
	"/favorites/list": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			favorites.FavoritesList,
		},
	},

	"/todo/add": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			todo.AddTodo,
		},
	},
	"/todo/update": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			todo.UpdateTodo,
		},
	},
	"/todo/delete": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			todo.DeleteTodo,
		},
	},
	"/todo/list": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			todo.TodoList,
		},
	},
}
