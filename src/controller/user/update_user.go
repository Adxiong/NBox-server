/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-09-02 22:25:26
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-02 23:57:38
 */
package user

import (
	"fmt"
	"log"
	user_service "nbox/src/model/service/user"

	"github.com/gin-gonic/gin"
)

type UpdateUserParams struct {
	user_service.UpdateUserParams
}

func UpdateUser(ctx *gin.Context) {

	params, err := checkParams(ctx)
	if err != nil {
		ctx.JSON(200, "msg: params is invalid")
	}

	userService := user_service.NewUserService()
	updateParams := user_service.UpdateUserParams{
		Username: params.Username,
		Password: params.Password,
	}
	_, errResp := userService.UpdateByID(ctx, 11, &updateParams)
	if errResp != nil {
		ctx.JSON(200, "msg:'update failed'")
	}
	ctx.JSON(200, "msg:'ok'")

}

func checkParams(ctx *gin.Context) (*UpdateUserParams, error) {
	params := &UpdateUserParams{}

	ctx.ShouldBindJSON(params)

	if params.Password == "" && params.Username == "" {
		log.Println("params is null")
		return params, fmt.Errorf("params is null")
	}

	return params, nil
}
