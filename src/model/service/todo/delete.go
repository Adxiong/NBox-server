/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 23:44:31
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-18 22:55:37
 */
package todo_service

import (
	"fmt"

	"golang.org/x/net/context"
)

func (*Todo) DeleteTodoByID(ctx context.Context, id uint64) (int64, error) {
	var rowsAffected int64
	var err error
	rowsAffected, err = deleteByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		return rowsAffected, err
	}
	return rowsAffected, err
}
