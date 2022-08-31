/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 09:58:00
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-31 23:11:48
 */
package main

import (
	"context"
	"nbox/src/bootstrap"
)

func main() {

	ctx, cancel := context.WithCancel((context.Background()))
	defer cancel()

	bootstrap.Start(ctx)
}
