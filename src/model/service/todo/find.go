/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 23:43:56
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-18 22:59:08
 */
package todo_service

import (
	"context"
	"fmt"
)

func (*Todo) FindTodoByUID(ctx context.Context, uid uint64) (*[]Todo, error) {
	var result *[]Todo = &[]Todo{}
	var err error
	result, err = findByUID(ctx, uid)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, err
}
