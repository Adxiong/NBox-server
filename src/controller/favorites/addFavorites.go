/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-05 22:44:13
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-05 23:24:21
 */
package favorites

import (
	"fmt"
	"log"
	favorites_service "nbox/src/model/service/favorites"

	"github.com/gin-gonic/gin"
)

type AddFavoritesParam struct {
	Title   string `json:"title" binding:"required"`
	Addr    string `json:"addr" binding:"required"`
	GroupID uint64 `json:"group_id" binding:"required"`
}

func AddFavorites(ctx *gin.Context) {
	param, errParam := checkAddFavoritesParams(ctx)
	if errParam != nil {
		log.Println("error", errParam)
		ctx.JSON(200, "param is invalid")
	}

	addParam := &favorites_service.Favorites{
		Title:   param.Title,
		Addr:    param.Addr,
		GroupID: param.GroupID,
	}

	service := favorites_service.NewFavorites()

	favorite, errAdd := service.Add(ctx, addParam)
	if errAdd != nil {
		log.Println("error", errAdd)
		log.Println("param", addParam)
		ctx.JSON(200, "add failed")
	}
	ctx.JSON(200, gin.H{"msg": "success", "data": *favorite})
}
func checkAddFavoritesParams(ctx *gin.Context) (*AddFavoritesParam, error) {
	param := &AddFavoritesParam{}
	errParam := ctx.ShouldBindJSON(param)
	if errParam != nil {
		log.Println(errParam)
		return param, errParam
	}

	if param.GroupID < 0 {
		err := fmt.Errorf("param-group_id is invalid")
		return param, err
	}
	return param, nil
}
