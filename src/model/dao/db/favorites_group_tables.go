/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 14:17:24
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-08 00:01:35
 */
package db

import "gorm.io/gorm"

type FavoritesGroup struct {
	gorm.Model

	Title   string `gorm:"column:title; comment:'群组标题'" json:"title"`
	Creator uint64 `gorm:"column:creator; comment:'群组创建者" json:"creator"`
	BgColor string `gorm:"column:bg_color; comment:'卡片背景颜色'" json:"bg_color"`
	Deleted uint8  `gorm:"column:deleted;comment:'已删除 0否 1是';not null; default:0" json:"deleted"`
}

var FavoritesGroupColumn = struct {
	ID        string
	Title     string
	Creator   string
	BgColor   string
	Deleted   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Title:     "title",
	Creator:   "creator",
	BgColor:   "bg_color",
	Deleted:   "deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

func (FavoritesGroup) TableName() string {
	return "favorites_group"
}
