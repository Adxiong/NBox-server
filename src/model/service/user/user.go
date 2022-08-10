/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-09 23:16:08
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-10 23:35:38
 */
package user

type user interface {
	AddUser()
	DelUserByUID()
	UpdateUserByUID()
	FindUserByUID()
	FindAllUser()
}
