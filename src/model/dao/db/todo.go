/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-07 22:25:15
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-18 23:42:51
 */
package db

import (
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type todo interface {
	AddTodo()
	DelTodoByID()
	UpdateTodoByID()
	FindTodoByUID()
}

func NewTodo() *Todo {
	return new(Todo)
}

func (t *Todo) AddTodo(ctx context.Context) (*Todo, error) {
	result := NewTodo()
	err := GlobalDb.Table(t.TableName()).Create(t).Error
	if err != nil {
		fmt.Println("创建错误", err)
		return result, fmt.Errorf("创建错误")
	}
	result = t
	return result, err
}

func (t *Todo) DelTodoByID(ctx context.Context, id uint64) (int64, error) {
	var res int64
	result := GlobalDb.Model(&Todo{}).Where("id = ?", id).Update(TodoColumn.IsDel, 1)
	if result.Error != nil {
		fmt.Println("删除错误", result.Error)
		return res, fmt.Errorf("删除错误")
	}
	return result.RowsAffected, nil
}

func (t *Todo) UpdateTodoByID(ctx context.Context, id uint64, vals map[string]interface{}) (int64, error) {
	var res int64
	result := GlobalDb.Model(&Todo{}).Where("id = ?", id).Updates(vals)
	if result.Error != nil {
		fmt.Println("更新错误", result.Error)
		return res, fmt.Errorf("更新失败")
	}
	return result.RowsAffected, nil
}

func (t *Todo) FindTodoByUID(ctx context.Context, uid uint64) (*[]Todo, error) {
	result := &[]Todo{}

	where := map[string]interface{}{
		TodoColumn.Creator: uid,
	}

	err := GlobalDb.Table(t.TableName()).Select("*").Where(where).Find(result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return result, err
	}
	return result, nil
}
