/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-01 23:11:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-01 23:25:21
 */
package db

import (
	"fmt"

	"golang.org/x/net/context"
)

type user interface {
	AddUser()
	DelUser()
	UpdateUser()
	GetUserById()
}

func NewUser() *User {
	return new(User)
}

func (user *User) AddUser(ctx context.Context) (*User, error) {
	res := NewUser()
	err := GlobalDb.Table(user.TableName()).Create(user).Error
	if err != nil {
		fmt.Println("创建错误", err)
		return res, fmt.Errorf("创建错误", err)

	}
	res = user
	return res, nil
}
