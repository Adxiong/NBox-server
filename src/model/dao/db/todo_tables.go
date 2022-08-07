/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 14:24:07
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-07 22:28:11
 */
package db

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Content string `gorm:"column:content;comment:'todo内容'" json:"content"`
	Status  uint8  `gorm:"column:status;comment:'todo状态 1进行中 2已完成'" json:"status"`
	Creator uint64 `gorm:"creator;comment:'创建者id'" json:"creator"`
}

var TodoColumn = struct {
	ID        string
	Content   string
	Status    string
	Creator   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Content:   "content",
	Status:    "status",
	Creator:   "creator",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

func (Todo) TableName() string {
	return "todo"
}
