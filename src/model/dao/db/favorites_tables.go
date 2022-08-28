/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 14:11:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-28 18:25:03
 */
package db

import "time"

type Favorites struct {
	ID        uint64     `gorm:"primarykey; column:id;" json:"id"`
	Title     string     `gorm:"column:title; comment:'收藏标题'" json:"title"`
	Addr      string     `gorm:"column:addr; comment:'收藏地址'" json:"addr"`
	GroupID   uint64     `gorm:"column:group_id; comment:'群组id'" json:"group_id"`
	IsDel     uint8      `gorm:"column:is_del;comment:'已删除 0否 1是';not null; default:0" json:"is_del"`
	Version   uint64     `gorm:"column:version;comment:'乐观锁'" json:"version"`
	CreateID  uint64     `gorm:"column:create_id; comment:'创建者id'" json:"create_id"`
	CreatedAt *time.Time `gorm:"column:create_at; comment:'创建时间'" json:"create_at"`
	UpdatedAt *time.Time `gorm:"column:update_at; comment:'更新时间'" json:"update_at"`
}

type FavoritesList []Favorites

var FavoritesColumn = struct {
	ID        string
	Title     string
	Addr      string
	GroupID   string
	IsDel     string
	Version   string
	CreateID  string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Title:     "title",
	Addr:      "addr",
	GroupID:   "group_id",
	IsDel:     "is_del",
	Version:   "version",
	CreateID:  "create_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (Favorites) TableName() string {
	return "favorites"
}
