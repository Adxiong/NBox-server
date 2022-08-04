/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-04 22:50:15
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-04 23:17:14
 */
package utils

import (
	"runtime"
	"strings"
)

func GetPackageAndFunc() string {
	pc, _, _, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(pc).Name()
	arr := strings.Split(caller, ".")
	for i, v := range arr {
		arr[i] = strings.Title(v)
	}
	result := strings.Join(arr, "_")
	return result
}
