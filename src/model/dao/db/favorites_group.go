/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 22:25:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-08 00:50:46
 */
package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type favoritesGroup interface {
	Add()
	DeleteByID()
	UpdateByID()
	FindByUID()
}

func NewFavoritesGroup() *FavoritesGroup {
	return new(FavoritesGroup)
}

func (f *FavoritesGroup) Add(ctx context.Context) (*FavoritesGroup, error) {
	result := NewFavoritesGroup()

	err := GlobalDb.Table(f.TableName()).Create(f).Error
	if err != nil {
		fmt.Println("增加失败", err)
		return result, err
	}
	result = f
	return result, nil
}

func (f *FavoritesGroup) DeleteByID(ctx context.Context, id uint64) (int64, error) {
	var res int64
	result := GlobalDb.Model(&FavoritesGroup{}).Where("id = ?", id).Update(FavoritesGroupColumn.Deleted, 1)
	if result.Error != nil {
		fmt.Println("删除失败")
		return res, result.Error
	}
	return result.RowsAffected, nil
}

func (f *FavoritesGroup) UpdateByID(ctx context.Context, id uint64, val map[string]interface{}) (int64, error) {
	var res int64
	result := GlobalDb.Model(&FavoritesGroup{}).Where("id = ?", id).Updates(val)
	if result.Error != nil {
		fmt.Println("更新失败")
		return res, result.Error
	}
	return result.RowsAffected, nil
}

func (f *FavoritesGroup) FindByUID(ctx context.Context, uid uint64) (*[]FavoritesGroup, error) {
	result := make([]FavoritesGroup, 0)

	err := GlobalDb.Table(f.TableName()).Where("createID = ?", uid).Find(result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &result, err
	}
	return &result, nil

}
