/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-05 22:46:52
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-05 23:55:20
 */
package favoritesgroup

import (
	"log"
	favoritesGroup "nbox/src/model/service/favoritesGroup/base"

	"github.com/gin-gonic/gin"
)

type AddFavoritesGroupParam struct {
	Title string
}

func AddFavoritesGroup(ctx *gin.Context) {
	param, errParam := checkAddFavoritesGroup(ctx)
	if errParam != nil {
		log.Println("error", errParam)
		ctx.JSON(200, gin.H{"msg": "param is invalid"})
	}
	
service := favoritesGroup.NewFavoritesGroup()
	
}
func checkAddFavoritesGroup(ctx *gin.Context) (*AddFavoritesGroupParam, error) {
	param := &AddFavoritesGroupParam{}
	err := ctx.ShouldBindJSON(param)
	if err != nil {
		return param, err
	}
	return param, nil
}
