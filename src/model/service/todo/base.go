/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 15:13:03
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-15 00:39:47
 */
package todo_service

type TodoService interface {
}

type Todo struct {
	ID        string
	Content   string
	Status    string
	Creator   string
	Deleted   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
