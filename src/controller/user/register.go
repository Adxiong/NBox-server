/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-30 23:07:18
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-04 22:15:22
 */
package user

import (
	"fmt"
	user_service "nbox/src/model/service/user"

	"github.com/gin-gonic/gin"
)

type createParams struct {
	user_service.User
}

func Register(ctx *gin.Context) {
	params, err := checkAddUserParams(ctx)
	if err != nil {
		ctx.JSON(200, "msg:"+err.Error())
	}

	user := &user_service.User{
		Email:    params.Email,
		Password: params.Password,
	}
	ret, err := user.Add(ctx, user)
	if err != nil {
		ctx.JSON(200, "msg:"+err.Error())
	}

	if ret.UID == 0 {
		ctx.JSON(200, "msg:"+err.Error())
	}
	ctx.JSON(200, "msg:'ok'")
}

func checkAddUserParams(ctx *gin.Context) (*createParams, error) {
	params := &createParams{}

	err := ctx.ShouldBindQuery(params)
	if err != nil {
		return nil, fmt.Errorf("params is invalid")
	}
	return params, nil
}
