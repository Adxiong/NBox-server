/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-09 23:16:08
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-09 23:24:22
 */
package user

import (
	"time"
)

type user interface {
	AddUser()
	DelUserByUID()
	UpdateUserByUID()
	FindUserByUID()
	FindAllUser()
}
type User struct {
	ID        uint
	UID       uint64
	Password  uint64
	Nick      string
	Name      string
	Birthday  time.Time
	Sex       uint8
	Salt      string
	Deleted   uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (User) AddUser() {

}

func (User) DelUserByUID() {

}

func (User) UpdateUserByUID() {

}

func (User) FindUserByUID() {

}

func (User) FindAllUser() {

}
