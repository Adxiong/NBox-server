/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 00:12:05
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-15 00:40:07
 */
package user_service

import "context"

func (*User) FindUserByUID(ctx context.Context, uid uint64) (*User, error) {
	var user *User

	res, err := findUserByUID(ctx, uid)
	if err != nil {
		return user, err
	}

	user = res
	return user, err
}
