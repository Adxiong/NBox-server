/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-02 23:45:33
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-02 23:54:34
 */
package response

type RespData struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Toast  string      `json:"Toast"`
	Data   interface{} `json:"data"`
}
