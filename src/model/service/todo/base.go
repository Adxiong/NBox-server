/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 15:13:03
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-25 22:46:49
 */
package todo_service

import (
	"context"
	"fmt"
	"nbox/src/model/dao/db"
	"time"
)

type todoService interface {
	Add(context.Context, *Todo) (*Todo, error)
	UpdateByTID(context.Context, uint64, *UpdateTodoParams) (int64, error)
	FindByUID(context.Context, uint64) (*[]Todo, error)
	DeleteByTID(context.Context, uint64) (int64, error)
}

type Todo struct {
	ID        uint64
	TID       uint64
	Content   string
	Status    uint8
	Creator   uint64
	IsDel     uint8
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type TodoList []Todo
type UpdateTodoParams struct {
	Content string
	Status  uint8
}

func NewTodoService() *Todo {
	return &Todo{}
}

func (todo *Todo) Add(ctx context.Context, params *Todo) (*Todo, error) {
	dbTodo := &db.Todo{
		ID:      params.ID,
		Content: params.Content,
		Status:  params.Status,
		Creator: params.Creator,
		IsDel:   0,
		Version: 1,
	}
	err := dbTodo.Add(ctx)
	if err != nil {
		fmt.Printf("SERVICE_TODO_BASE_Add_Failed.\nerror:%s,params:%+v\n", err.Error(), dbTodo)
		return &Todo{}, err
	}
	result := &Todo{
		ID:        dbTodo.ID,
		TID:       dbTodo.TID,
		Content:   dbTodo.Content,
		Status:    dbTodo.Status,
		Creator:   dbTodo.Creator,
		IsDel:     dbTodo.IsDel,
		CreatedAt: dbTodo.CreatedAt,
		UpdatedAt: dbTodo.UpdatedAt,
	}
	return result, nil
}

func (todo *Todo) UpdateByTID(ctx context.Context, tid uint64, params *UpdateTodoParams) (int64, error) {
	var rowsAffected int64
	var err error

	val := map[string]interface{}{}

	if params.Content != "" {
		val[db.TodoColumn.Content] = params.Content
	}

	if params.Status >= 0 {
		val[db.TodoColumn.Status] = params.Status
	}

	if len(val) == 0 {
		fmt.Println("SERVICE_TODO_BASE_UpdateByTID_Params_Empty")
		return rowsAffected, err
	}
	todoDao := db.NewTodo()
	rowsAffected, err = todoDao.UpdateByTID(ctx, tid, val)
	if err != nil {
		fmt.Printf("SERVICE_TODO_BASE_UpdateByTID_Failed.\nerror:%s,tid:%d,params:%+v\n", err.Error(), tid, val)
		return rowsAffected, err
	}
	return rowsAffected, err
}

func (todo *Todo) DeleteByTID(ctx context.Context, tid uint64) (int64, error) {
	var rowsAffected int64
	var err error

	todoDao := db.NewTodo()
	rowsAffected, err = todoDao.DeleteByTID(ctx, tid)
	if err != nil {
		fmt.Printf("SERVICE_TODO_BASE_DeleteByTID_Failed.\nerror:%s,tid:%d\n", err.Error(), tid)
		return rowsAffected, err
	}
	return rowsAffected, err
}

func (todo *Todo) FindListByUID(ctx context.Context, uid uint64, page uint64, size uint64) (*TodoList, error) {
	if size == 0 {
		size = 20
	}
	todoDao := db.NewTodo()
	todoList, err := todoDao.FindListByUID(ctx, uid, page, size)
	if err != nil {
		fmt.Println(err)
		return &TodoList{}, err
	}
	result := make(TodoList, 0)

	for _, item := range *todoList {
		if item.IsDel != 0 {
			continue
		}
		todo := &Todo{
			ID:        item.ID,
			Content:   item.Content,
			Status:    item.Status,
			Creator:   item.Creator,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		result = append(result, *todo)
	}
	return &result, nil
}
