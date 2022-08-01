/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 12:12:02
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-01 23:52:51
 */
package db

import (
	"time"
)

type User struct {
	ID       uint `gorm:"column:id;type:bigint(20);autoIncrement:true;primaryKey;not null;" json:"-"`
	UId      uint64
	Password string
	Nick     string
	Name     string
	Birthday time.Time
	Sex      uint8
	Salt     string
}

type TransformedUser struct {
	ID       uint      `json:"id"`
	UId      uint64    `json:"uid"`
	Password string    `json:"password"`
	Nick     string    `json:"nick"`
	Birthday time.Time `json:"birthday"`
	Sex      uint8     `json:"sex"`
	Salt     string    `json:"salt"`
}

func (u *User) TableName() string {
	return "user"
}
