/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 12:12:02
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-07-31 23:04:07
 */
package dao

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20)"`
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
