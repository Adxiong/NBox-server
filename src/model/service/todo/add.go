/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 23:43:48
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-18 22:56:59
 */
package todo_service

import (
	"context"
	"fmt"
)

func (*Todo) AddTodo(ctx context.Context, params *Todo) (*Todo, error) {
	var result *Todo = &Todo{}
	var err error
	result, err = add(ctx, params)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, err
}
