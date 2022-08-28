/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 15:13:03
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-28 22:53:09
 */
package todo_service

import (
	"context"
	"fmt"
	"nbox/src/model/dao/db"
	"time"
)

type TodoService interface {
	AddTodo(context.Context, *Todo) (*Todo, error)
	UpdateTodoByID(context.Context, uint64, *UpdateTodoParams) (int64, error)
	FindTodoByUID(context.Context, uint64) (*[]Todo, error)
	DeleteByID(context.Context, uint64) (int64, error)
}

type Todo struct {
	ID        uint64
	Content   string
	Status    uint8
	Creator   uint64
	IsDel     uint8
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type UpdateTodoParams struct {
	Content string
	Status  string
	Deleted string
}

func add(ctx context.Context, params *Todo) (*Todo, error) {
	dbTodo := &db.Todo{
		ID:      params.ID,
		Content: params.Content,
		Status:  params.Status,
		Creator: params.Creator,
		IsDel:   params.IsDel,
	}
	err := dbTodo.AddTodo(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return &Todo{}, err
	}
	result := &Todo{
		ID:        dbTodo.ID,
		Content:   dbTodo.Content,
		Status:    dbTodo.Status,
		Creator:   dbTodo.Creator,
		IsDel:     dbTodo.IsDel,
		CreatedAt: dbTodo.CreatedAt,
		UpdatedAt: dbTodo.UpdatedAt,
	}
	return result, nil
}

func updateByID(ctx context.Context, id uint64, params *UpdateTodoParams) (int64, error) {
	var rowsAffected int64
	var err error

	dbVals := map[string]interface{}{}

	if params.Content != "" {
		dbVals[db.TodoColumn.Content] = params.Content
	}

	if params.Status != "" {
		dbVals[db.TodoColumn.Status] = params.Status
	}

	if params.Deleted != "" {
		dbVals[db.TodoColumn.IsDel] = params.Deleted
	}

	todoDao := db.NewTodo()
	rowsAffected, err = todoDao.UpdateTodoByID(ctx, id, dbVals)
	if err != nil {
		fmt.Println(err)
		return rowsAffected, err
	}
	return rowsAffected, err
}

func deleteByID(ctx context.Context, id uint64) (int64, error) {
	var rowsAffected int64
	var err error

	todoDao := db.NewTodo()
	rowsAffected, err = todoDao.DelTodoByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		return rowsAffected, err
	}
	return rowsAffected, err
}

func findByUID(ctx context.Context, uid uint64) (*[]Todo, error) {
	var result = []Todo{}

	todoDao := db.NewTodo()
	todoList, err := todoDao.FindTodoByUID(ctx, uid)
	if err != nil {
		fmt.Println(err)
		return &result, err
	}

	for _, item := range *todoList {
		todo := &Todo{
			ID:        item.ID,
			Content:   item.Content,
			Status:    item.Status,
			Creator:   item.Creator,
			IsDel:     item.IsDel,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		result = append(result, *todo)
	}
	return &result, nil
}
