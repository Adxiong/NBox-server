/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 00:11:33
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-15 00:40:01
 */
package user_service

import "context"

func (*User) AddUser(ctx context.Context, param *User) (*User, error) {
	var user *User
	res, err := add(ctx, param)
	if err != nil {
		return user, err
	}
	user = res
	return user, nil
}
