/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-01 23:11:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-10 23:41:29
 */
package db

import (
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type user interface {
	Add()
	UpdateByID()
	FindByUID()
}

func NewUser() *User {
	return new(User)
}

// AddUser 新增用户记录
func (user *User) Add(ctx context.Context) (*User, error) {
	res := NewUser()
	err := GlobalDb.Table(user.TableName()).Create(user).Error
	if err != nil {
		fmt.Println("创建错误", err)
		return res, fmt.Errorf("创建错误")
	}
	res = user
	return res, nil
}

// UpdateByID 更新用户信息
func (user *User) UpdateByID(ctx context.Context, id uint64, vals map[string]interface{}) (int64, error) {
	var res int64
	result := GlobalDb.Model(&User{}).Where("id = ?", id).Updates(vals)
	if result.Error != nil {
		fmt.Println("更新错误", result.Error)
		return res, fmt.Errorf("更新错误")
	}
	return result.RowsAffected, nil
}

// FindByUID 根据用户id查询用户信息
func (user *User) FindByUID(ctx context.Context, uid uint64) (*User, error) {
	res := NewUser()

	where := map[string]interface{}{
		UserColumn.UID: uid,
	}
	err := GlobalDb.Table(user.TableName()).Select("*").Where(where).First(res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return res, err
	}
	return res, nil
}
