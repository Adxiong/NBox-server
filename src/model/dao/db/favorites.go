/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 22:25:34
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-28 18:34:17
 */
package db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type favorites interface {
	Add()
	UpdateByID()
	DeleteByID()
	FindListByGID()
}

func NewFavorites() *Favorites {
	return new(Favorites)
}

// Add 增加favorites记录
func (f *Favorites) Add(ctx context.Context) (*Favorites, error) {
	result := NewFavorites()

	err := GlobalDb.Table(f.TableName()).Create(f).Error
	if err != nil {
		log.Println("添加失败", err)
		return result, fmt.Errorf("添加失败")
	}
	result = f
	return result, nil
}

func (f *Favorites) UpdateTodoByID(ctx context.Context, id uint64, vals map[string]interface{}) (int64, error) {
	var res int64
	result := GlobalDb.Model(&Favorites{}).Where("id = ?", id).Updates(vals)
	if result.Error != nil {
		log.Println("更新失败", result.Error)
		return res, fmt.Errorf("更新失败")
	}
	return result.RowsAffected, nil
}

func (f *Favorites) DeleteByID(ctx context.Context, id uint64) (int64, error) {
	var res int64

	result := GlobalDb.Model(&Favorites{}).Where("id = ?", id).Update(FavoritesColumn.IsDel, 1)
	if result.Error != nil {
		log.Println("删除失败", result.Error)
		return res, fmt.Errorf("删除失败")
	}
	return result.RowsAffected, nil
}

func (f *Favorites) FindListByGID(ctx context.Context, gid uint64) (*FavoritesList, error) {
	var res = make(FavoritesList, 0)
	err := GlobalDb.Table(f.TableName()).Where("gid = ?", gid).Find(res).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("查询失败", err)
		return &res, fmt.Errorf("查询失败")
	}
	return &res, nil
}
