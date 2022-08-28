/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 22:25:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-28 23:08:24
 */
package db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type favoritesGroup interface {
	Add()
	DeleteByGID()
	UpdateByGID()
	FindListByUID()
}

func NewFavoritesGroup() *FavoritesGroup {
	return new(FavoritesGroup)
}

func (f *FavoritesGroup) Add(ctx context.Context) error {
	err := GlobalDb.Table(f.TableName()).Create(f).Error
	if err != nil {
		log.Println("添加失败", err)
		return fmt.Errorf("添加失败")
	}
	return nil
}

func (f *FavoritesGroup) DeleteByGID(ctx context.Context, gid uint64) (int64, error) {
	var res int64
	result := GlobalDb.Model(&FavoritesGroup{}).Where("gid = ?", gid).Update(FavoritesGroupColumn.IsDel, 1)
	if result.Error != nil {
		log.Println("删除失败", result.Error)
		return res, fmt.Errorf("删除失败")
	}
	return result.RowsAffected, nil
}

func (f *FavoritesGroup) UpdateByGID(ctx context.Context, gid uint64, val map[string]interface{}) (int64, error) {
	var res int64
	result := GlobalDb.Model(&FavoritesGroup{}).Where("gid = ?", gid).Updates(val)
	if result.Error != nil {
		log.Println("更新失败", result.Error)
		return res, fmt.Errorf("更新失败")
	}
	return result.RowsAffected, nil
}

func (f *FavoritesGroup) FindListByUID(ctx context.Context, uid uint64) (*FavoritesGroupList, error) {
	var res = make(FavoritesGroupList, 0)
	err := GlobalDb.Table(f.TableName()).Where("createID = ?", uid).Find(res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("查询失败", err)
		return &res, fmt.Errorf("查询失败")
	}
	return &res, nil
}
