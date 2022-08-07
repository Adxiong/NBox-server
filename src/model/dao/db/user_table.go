/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-07-31 12:12:02
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-08 00:00:53
 */
package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID      uint64     `gorm:"column:uid;comment:'用户唯一id'" json:"uid"`
	Password string     `gorm:"column:password;comment:'密码'" json:"password"`
	Nick     string     `gorm:"column:nick;comment:'昵称'" json:"nick"`
	Name     string     `gorm:"column:name;comment:'姓名'" json:"name"`
	Birthday *time.Time `gorm:"column:birthday;comment:'生日'" json:"birthday"`
	Sex      uint8      `gorm:"column:sex;comment:'性别 1男 2女'" json:"sex"`
	Salt     string     `gorm:"column:sale;comment:'盐 密码加盐'" json:"salt"`
	Deleted  uint8      `gorm:"column:deleted;comment:'已删除 0否 1是';not null; default:0" json:"deleted"`
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
	Deleted   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	UID:       "uid",
	Password:  "password",
	Nick:      "nick",
	Name:      "name",
	Birthday:  "birthday",
	Sex:       "sex",
	Salt:      "salt",
	Deleted:   "deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_ad",
}

func (user *User) TableName() string {
	return "user"
}
