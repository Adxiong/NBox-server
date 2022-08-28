/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 14:24:07
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-28 18:26:55
 */
package db

import "time"

type Todo struct {
	ID        uint64     `gorm:"primarykey;column:id" json:"id"`
	Content   string     `gorm:"column:content;comment:'todo内容'" json:"content"`
	Status    uint8      `gorm:"column:status;comment:'todo状态 1进行中 2已完成'" json:"status"`
	Creator   uint64     `gorm:"column:creator;comment:'创建者id'" json:"creator"`
	IsDel     uint8      `gorm:"column:is_del;comment:'已删除 0否 1是';not null; default:0" json:"is_del"`
	Version   uint64     `gorm:"column:version; comment:'乐观锁'" json:"version"`
	CreatedAt *time.Time `gorm:"column:create_at; comment:'创建时间'" json:"create_at"`
	UpdatedAt *time.Time `grom:"column:update_at; comment:'更新时间'" json:"update_at"`
}

type TodoList []Todo

var TodoColumn = struct {
	ID        string
	Content   string
	Status    string
	Creator   string
	IsDel     string
	Version   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Content:   "content",
	Status:    "status",
	Creator:   "creator",
	IsDel:     "is_del",
	Version:   "version",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (Todo) TableName() string {
	return "todo"
}
