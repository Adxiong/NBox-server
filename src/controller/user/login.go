/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-09-03 00:03:09
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-25 22:24:00
 */
package user

import (
	"log"
	user_service "nbox/src/model/service/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

func Login(ctx *gin.Context) {
	params, err := checkLoginParams(ctx)
	if err != nil {
		ctx.JSON(200, "msg: params is invalid")
		return
	}
	session := sessions.Default(ctx)

	if session.Get("user:"+params.Email) != nil {
		ctx.JSON(200, gin.H{"msg": "success", "user1": session.Get("user:" + params.Email)})
		return
	}

	service := user_service.NewUserService()
	userInfo, errUser := service.FindByEmail(ctx, params.Email)
	if errUser != nil {
		log.Println("error", errUser)
		ctx.JSON(200, "msg: email is not exist")
	}
	// todo 此处密码应该解密
	if userInfo.Password != params.Password {
		ctx.JSON(200, gin.H{"msg": "password error"})
		return
	}

	val := map[string]string{
		"username": userInfo.Username,
		"email":    userInfo.Email,
		// todo uid加密返回
		"uid": string(userInfo.UID),
	}
	session.Set("user:"+params.Email, val)
	session.Save()
	ctx.JSON(200, gin.H{"msg": "login success", "user": val})
}

func checkLoginParams(ctx *gin.Context) (*LoginParams, error) {
	params := &LoginParams{}

	err := ctx.ShouldBind(params)
	if err != nil {
		log.Println("error", err)
		return params, err
	}
	// todo 校验邮箱格式
	return params, nil
}
