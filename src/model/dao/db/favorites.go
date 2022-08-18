/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 22:25:34
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-18 23:37:03
 */
package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type favorites interface {
	Add()
	UpdateByID()
	DeleteByID()
	FindByGID()
}

func NewFavorites() *Favorites {
	return new(Favorites)
}

// Add 增加favorites记录
func (f *Favorites) Add(ctx context.Context) (*Favorites, error) {
	result := NewFavorites()

	err := GlobalDb.Table(f.TableName()).Create(f).Error
	if err != nil {
		fmt.Println("增加失败", err)
		return result, fmt.Errorf("增加失败")
	}
	result = f
	return result, nil
}

func (f *Favorites) UpdateTodoByID(ctx context.Context, id uint64, vals map[string]interface{}) (int64, error) {
	var res int64
	result := GlobalDb.Model(&Favorites{}).Where("id = ?", id).Updates(vals)
	if result.Error != nil {
		fmt.Println("更新", result.Error)
		return res, fmt.Errorf("更新失败")
	}
	return result.RowsAffected, nil
}

func (f *Favorites) DeleteByID(ctx context.Context, id uint64) (int64, error) {
	var res int64

	result := GlobalDb.Model(&Favorites{}).Where("id = ?", id).Update(FavoritesColumn.IsDel, 1)
	if result.Error != nil {
		fmt.Println("删除失败", result.Error)
		return res, fmt.Errorf("删除失败")
	}
	return result.RowsAffected, nil
}

func (f *Favorites) FindByGID(ctx context.Context, gid uint64) ([]Favorites, error) {
	result := make([]Favorites, 0)

	err := GlobalDb.Table(f.TableName()).Where("gid = ?", gid).Find(result).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return result, err
	}
	return result, err
}
