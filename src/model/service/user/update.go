/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 00:11:50
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-14 00:27:14
 */
package user

import "context"

func (*User) UpdateByID(ctx context.Context, id uint64, params *UpdateUserParams) (int64, error) {
	var rowsAffected int64

	res, err := updateByID(ctx, id, params)
	if err != nil {
		return rowsAffected, err
	}

	rowsAffected = res
	return rowsAffected, nil
}
