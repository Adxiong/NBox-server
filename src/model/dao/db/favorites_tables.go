/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 14:11:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-08 00:01:01
 */
package db

import "gorm.io/gorm"

type Favorites struct {
	gorm.Model

	Title   string `gorm:"column:title; comment:'收藏标题'" json:"title"`
	Addr    string `gorm:"column:addr; comment:'收藏地址'" json:"addr"`
	GroupID uint64 `gorm:"column:group_id; comment:'群组id'" json:"group_id"`
	Deleted uint8  `gorm:"column:deleted;comment:'已删除 0否 1是';not null; default:0" json:"deleted"`
}

var FavoritesColumn = struct {
	ID        string
	Title     string
	Addr      string
	GroupID   string
	Deleted   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Title:     "title",
	Addr:      "addr",
	GroupID:   "group_id",
	Deleted:   "deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

func (Favorites) TableName() string {
	return "favorites"
}
