/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-09 23:16:19
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-04 22:15:02
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
	FindByEmail()
}

type User struct {
	ID        uint64     `json:"id" `
	UID       uint64     `json:"uid"`
	Password  string     `json:"password"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"create_at"`
	UpdatedAt *time.Time `json:"update_at"`
	IsDel     uint8      `json:"-"`
	Version   uint64     `json:"-"`
}

type UpdateUserParams struct {
	Password string
	Username string
}

func NewUserService() *User {
	return new(User)
}

func (*User) Add(ctx context.Context, params *User) (*User, error) {
	dbUser := &db.User{
		ID:       params.ID,
		UID:      params.UID,
		Password: params.Password,
		Username: params.Username,
		Email:    params.Email,
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
		Username: user.Username,
		Email:    user.Email,
	}
	return result, nil
}

func (*User) UpdateByID(ctx context.Context, id uint64, params *UpdateUserParams) (int64, error) {
	var rowsAffected int64
	var err error = nil

	val := map[string]interface{}{}

	if params.Username != "" {
		val[db.UserColumn.Username] = params.Username
	}

	if params.Password != "" {
		val[db.UserColumn.Password] = params.Password
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

func (*User) FindByUID(ctx context.Context, uid uint64) (*User, error) {
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
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
	return resUser, nil
}

func (*User) FindByEmail(ctx context.Context, email string) (*User, error) {
	resUser := new(User)
	var err error = nil

	dbDao := db.NewUser()
	dbUser, err := dbDao.FindByEmail(ctx, email)

	if err != nil {
		fmt.Printf("SERVICE_USER_BASE_FindUserByUID_Failed,[error] %s\n", err.Error())
		return resUser, err
	}
	resUser = &User{
		ID:        dbUser.ID,
		UID:       dbUser.UID,
		Password:  dbUser.Password,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
	return resUser, nil
}
