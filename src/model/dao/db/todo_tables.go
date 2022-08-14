/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 14:24:07
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-14 23:43:27
 */
package db

import "time"

type Todo struct {
	ID        uint64     `gorm:"primarykey;column:id" json:"id"`
	CreatedAt *time.Time `gorm:"column:create_at; comment:'创建时间'" json:"create_at"`
	UpdatedAt *time.Time `grom:"column:update_at; comment:'更新时间'" json:"update_at"`
	Content   string     `gorm:"column:content;comment:'todo内容'" json:"content"`
	Status    uint8      `gorm:"column:status;comment:'todo状态 1进行中 2已完成'" json:"status"`
	Creator   uint64     `gorm:"column:creator;comment:'创建者id'" json:"creator"`
	Deleted   uint8      `gorm:"column:deleted;comment:'已删除 0否 1是';not null; default:0" json:"deleted"`
}

var TodoColumn = struct {
	ID        string
	Content   string
	Status    string
	Creator   string
	Deleted   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Content:   "content",
	Status:    "status",
	Creator:   "creator",
	Deleted:   "deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

func (Todo) TableName() string {
	return "todo"
}
