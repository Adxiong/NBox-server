/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 12:12:02
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-10 23:40:41
 */
package db

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primarykey"`
	UID       uint64     `gorm:"column:uid;comment:'用户唯一id'" json:"uid"`
	Password  string     `gorm:"column:password;comment:'密码'" json:"password"`
	Nick      string     `gorm:"column:nick;comment:'昵称'" json:"nick"`
	Name      string     `gorm:"column:name;comment:'姓名'" json:"name"`
	Birthday  *time.Time `gorm:"column:birthday;comment:'生日'" json:"birthday"`
	Sex       uint8      `gorm:"column:sex;comment:'性别 1男 2女'" json:"sex"`
	Salt      string     `gorm:"column:sale;comment:'盐 密码加盐'" json:"salt"`
	CreatedAt time.Time  `gorm:"column:create_at;comment:'创建时间'" json:"create_at"`
	UpdatedAt time.Time  `gorm:"column:update_at;comment:'更新时间'" json:"update_at"`
}

var UserColumn = struct {
	ID        string
	UID       string
	Password  string
	Nick      string
	Name      string
	Birthday  string
	Sex       string
	Salt      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UID:       "uid",
	Password:  "password",
	Nick:      "nick",
	Name:      "name",
	Birthday:  "birthday",
	Sex:       "sex",
	Salt:      "salt",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (user *User) TableName() string {
	return "user"
}
