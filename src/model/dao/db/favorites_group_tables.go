/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 14:17:24
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-18 23:40:14
 */
package db

import "time"

type FavoritesGroup struct {
	ID        uint64     `gorm:"primarykey; column:id" json:"id"`
	Title     string     `gorm:"column:title; comment:'群组标题'" json:"title"`
	Creator   uint64     `gorm:"column:creator; comment:'群组创建者" json:"creator"`
	BgColor   string     `gorm:"column:bg_color; comment:'卡片背景颜色'" json:"bg_color"`
	IsDel     uint8      `gorm:"column:is_del; comment:'已删除 0否 1是';not null; default:0" json:"is_del"`
	Version   uint64     `gorm:"column:version; comment:'乐观锁'" json:"version"`
	CreatedAt *time.Time `gorm:"column:create_at; comment:'创建时间'" json:"create_at"`
	UpdatedAt *time.Time `gorm:"column:update_at; comment:'更新时间'" json:"update_at"`
}

var FavoritesGroupColumn = struct {
	ID        string
	Title     string
	Creator   string
	BgColor   string
	IsDel     string
	Version   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Title:     "title",
	Creator:   "creator",
	BgColor:   "bg_color",
	IsDel:     "is_del",
	Version:   "version",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (FavoritesGroup) TableName() string {
	return "favorites_group"
}
