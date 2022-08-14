/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-09 23:16:19
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-14 15:12:07
 */
package user

import (
	"context"
	"fmt"
	"nbox/src/model/dao/db"
	"time"
)

type userService interface {
	AddUser() (*User, error)
	UpdateUserByUID()
	FindUserByUID()
}

type User struct {
	ID        uint
	UID       uint64
	Password  string
	Nick      string
	Name      string
	Birthday  *time.Time
	Sex       uint8
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
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

func add(ctx context.Context, params *User) (*User, error) {

	user := &db.User{
		ID:       params.ID,
		UID:      params.UID,
		Password: params.Password,
		Name:     params.Name,
		Nick:     params.Nick,
		Birthday: params.Birthday,
		Sex:      params.Sex,
		Salt:     params.Salt,
	}

	dbUser, err := user.Add(ctx)
	if err != nil {
		fmt.Println("SERVICE_USER_Add_Failed", err)
		return &User{}, err
	}
	result := &User{
		ID:       dbUser.ID,
		UID:      dbUser.UID,
		Name:     dbUser.Name,
		Password: dbUser.Password,
		Nick:     dbUser.Nick,
		Birthday: dbUser.Birthday,
		Sex:      dbUser.Sex,
		Salt:     dbUser.Salt,
	}
	return result, nil
}

func updateByID(ctx context.Context, id uint64, params *UpdateUserParams) (int64, error) {
	var rowsAffected int64
	var err error = nil

	dbVals := map[string]interface{}{}

	if params.Nick != nil {
		dbVals[db.UserColumn.Nick] = params.Nick
	}

	if params.Name != nil {
		dbVals[db.UserColumn.Name] = params.Name
	}

	if params.Password != nil {
		dbVals[db.UserColumn.Password] = params.Password
	}

	if params.Birthday != nil {
		dbVals[db.UserColumn.Birthday] = params.Birthday
	}

	if params.Sex != nil {
		dbVals[db.UserColumn.Sex] = params.Sex
	}

	if params.Salt != nil {
		dbVals[db.UserColumn.Salt] = params.Salt
	}

	if len(dbVals) == 0 {
		fmt.Println("SERVICE_USER_BASE_updateByID_Params_Empty")
		return rowsAffected, err
	}

	dbDao := db.NewUser()

	rowsAffected, err = dbDao.UpdateByID(ctx, id, dbVals)

	if err != nil {
		fmt.Printf("SERVICE_USER_BASE_updateByID_Failed,[error] %s\n", err.Error())
		return rowsAffected, err
	}
	return rowsAffected, err
}

func findUserByUID(ctx context.Context, uid uint64) (*User, error) {
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
