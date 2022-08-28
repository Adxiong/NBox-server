/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-04 22:50:15
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-04 23:35:36
 */
package utils

import (
	"runtime"
	"strings"
)

const (
	LOG_WARNING = "Warning"
	LOG_NOTICE  = "Notice"
	LOG_ERROR   = "Error"
	LOG_TRACE   = "Trace"
	LOG_FATAL   = "Fatal"
	LOG_OUTPUT  = "Output"
)

func GenLogMessage(logType string, actions string) string {
	str := logType + "_" + GetPackageAndFunc() + "_" + actions
	return str
}

func GetPackageAndFunc() string {
	pc, _, _, _ := runtime.Caller(2)
	caller := runtime.FuncForPC(pc).Name()
	arr := strings.Split(caller, ".")
	for i, v := range arr {
		arr[i] = strings.Title(v)
	}
	result := strings.Join(arr, "_")
	return result
}
