/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-09 23:16:19
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-30 00:28:40
 */
package user_service

import (
	"context"
	"fmt"
	"nbox/src/model/dao/db"
	"time"
)

type userService interface {
	Add()
	UpdateByUID()
	FindByUID()
}

type User struct {
	ID        uint64
	UID       uint64
	Password  string
	Nick      string
	Name      string
	Birthday  *time.Time
	Sex       uint8
	Salt      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	IsDel     uint8
	Version   uint64
}

type UpdateUserParams struct {
	Password *string
	Nick     *string
	Name     *string
	Birthday *time.Time
	Sex      *uint8
	Salt     *string
}

func NewUserService() *User {
	return new(User)
}

func (*User) Add(ctx context.Context, params *User) (*User, error) {

	dbUser := &db.User{
		ID:       params.ID,
		UID:      params.UID,
		Password: params.Password,
		Name:     params.Name,
		Nick:     params.Nick,
		Birthday: params.Birthday,
		Sex:      params.Sex,
		Salt:     params.Salt,
		IsDel:    0,
		Version:  1,
	}

	user, err := dbUser.Add(ctx)
	if err != nil {
		fmt.Println("SERVICE_USER_Add_Failed", err)
		return &User{}, err
	}
	result := &User{
		ID:       user.ID,
		UID:      user.UID,
		Name:     user.Name,
		Password: user.Password,
		Nick:     user.Nick,
		Birthday: user.Birthday,
		Sex:      user.Sex,
		Salt:     user.Salt,
	}
	return result, nil
}

func (*User) UpdateByID(ctx context.Context, id uint64, params *UpdateUserParams) (int64, error) {
	var rowsAffected int64
	var err error = nil

	val := map[string]interface{}{}

	if params.Nick != nil {
		val[db.UserColumn.Nick] = params.Nick
	}

	if params.Name != nil {
		val[db.UserColumn.Name] = params.Name
	}

	if params.Password != nil {
		val[db.UserColumn.Password] = params.Password
	}

	if params.Birthday != nil {
		val[db.UserColumn.Birthday] = params.Birthday
	}

	if params.Sex != nil {
		val[db.UserColumn.Sex] = params.Sex
	}

	if params.Salt != nil {
		val[db.UserColumn.Salt] = params.Salt
	}

	if len(val) == 0 {
		fmt.Println("SERVICE_USER_BASE_updateByID_Params_Empty")
		return rowsAffected, err
	}

	dbDao := db.NewUser()

	rowsAffected, err = dbDao.UpdateByID(ctx, id, val)

	if err != nil {
		fmt.Printf("SERVICE_USER_BASE_updateByID_Failed,[error] %s\n", err.Error())
		return rowsAffected, err
	}
	return rowsAffected, err
}

func (*User) FindListByUID(ctx context.Context, uid uint64) (*User, error) {
	resUser := new(User)
	var err error = nil

	dbDao := db.NewUser()
	dbUser, err := dbDao.FindByUID(ctx, uid)

	if err != nil {
		fmt.Printf("SERVICE_USER_BASE_FindUserByUID_Failed,[error] %s\n", err.Error())
		return resUser, err
	}
	resUser = &User{
		ID:        dbUser.ID,
		UID:       dbUser.UID,
		Password:  dbUser.Password,
		Nick:      dbUser.Nick,
		Name:      dbUser.Name,
		Birthday:  dbUser.Birthday,
		Sex:       dbUser.Sex,
		Salt:      dbUser.Salt,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
	return resUser, nil
}
