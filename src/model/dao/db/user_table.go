/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 12:12:02
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-02 00:29:01
 */
package db

import (
	"time"
)

type User struct {
	ID       int       `gorm:"autoIncrement; primaryKey; column:id; type:bigint(20); not null;" json:"-"`
	UId      uint64    `gorm:"column:uid"`
	Password string    `gorm:"column:password"`
	Nick     string    `gorm:"column:nick"`
	Name     string    `gorm:"column:name"`
	Birthday time.Time `gorm:"column:birthday"`
	Sex      uint8     `gorm:"column:sex"`
	Salt     string    `gorm:"column:sale"`
}

type TransformedUser struct {
	ID       uint      `json:"id"`
	UId      uint64    `json:"uid"`
	Password string    `json:"password"`
	Nick     string    `json:"nick"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Sex      uint8     `json:"sex"`
	Salt     string    `json:"salt"`
}

func (u *User) TableName() string {
	return "user"
}
