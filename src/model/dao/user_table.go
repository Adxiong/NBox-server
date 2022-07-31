/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 12:12:02
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-07-31 15:21:56
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
	nick     string
	name     string
	birthday time.Time
	sex      uint8
	salt     string
}

type transformedUser struct {
	ID       uint      `json:"id"`
	UId      uint64    `json:"uid"`
	Password string    `json:"password"`
	nick     string    `json:"nick"`
	birthday time.Time `json:"birthday"`
	sex      uint8     `json:"sex"`
	salt     string    `json:"salt`
}

func (u *User) TableName() string {
	return "user"
}
