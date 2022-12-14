/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 22:25:15
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-25 22:47:21
 */
package db

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type todo interface {
	Add()
	DeleteByID()
	UpdateByID()
	FindListByUID()
}

func NewTodo() *Todo {
	return new(Todo)
}

func (t *Todo) Add(ctx context.Context) error {
	err := GlobalDb.Table(t.TableName()).Create(t).Error
	if err != nil {
		log.Println("添加失败", err)
		return fmt.Errorf("添加失败")
	}
	return err
}

func (t *Todo) DeleteByTID(ctx context.Context, tid uint64) (int64, error) {
	var res int64
	result := GlobalDb.Model(&Todo{}).Where("tid = ?", tid).Update(TodoColumn.IsDel, 1)
	if result.Error != nil {
		log.Println("删除失败", result.Error)
		return res, fmt.Errorf("删除失败")
	}
	return result.RowsAffected, nil
}

func (t *Todo) UpdateByTID(ctx context.Context, tid uint64, vals map[string]interface{}) (int64, error) {
	var res int64
	result := GlobalDb.Model(&Todo{}).Where("tid = ?", tid).Updates(vals)
	if result.Error != nil {
		log.Println("更新失败", result.Error)
		return res, fmt.Errorf("更新失败")
	}
	return result.RowsAffected, nil
}

func (t *Todo) FindListByUID(ctx context.Context, uid uint64, page uint64, size uint64) (*TodoList, error) {
	var result = make(TodoList, 0)

	where := map[string]interface{}{
		TodoColumn.Creator: uid,
	}
	offset := page * size
	err := GlobalDb.Table(t.TableName()).Select("*").Where(where).Offset(int(offset)).Limit(int(size)).Find(result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("查询失败", err)
		return &result, fmt.Errorf("查询失败")
	}
	return &result, nil
}
