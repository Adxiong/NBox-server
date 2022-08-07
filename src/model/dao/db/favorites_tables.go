/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 14:11:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-07 22:30:38
 */
package db

import "gorm.io/gorm"

type Favorites struct {
	gorm.Model

	Title   string `gorm:"column:title; comment:'收藏标题'" json:"title"`
	Addr    string `gorm:"column:addr; comment:'收藏地址'" json:"addr"`
	GroupID uint64 `gorm:"column:group_id; comment:'群组id'" json:"group_id"`
}

var FavoritesColumn = struct {
	ID        string
	Title     string
	Addr      string
	GroupID   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Title:     "title",
	Addr:      "addr",
	GroupID:   "group_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

func (Favorites) TableName() string {
	return "favorites"
}
