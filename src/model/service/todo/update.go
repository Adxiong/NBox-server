/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 23:44:05
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-18 23:01:44
 */
package todo_service

import (
	"context"
	"fmt"
)

func (*Todo) UpdateTodoByID(ctx context.Context, id uint64, params *UpdateTodoParams) (int64, error) {
	var rowsAffected int64
	var err error
	rowsAffected, err = updateByID(ctx, id, params)
	if err != nil {
		fmt.Println(err)
		return rowsAffected, err
	}
	return rowsAffected, err
}
