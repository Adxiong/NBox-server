/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 15:13:03
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-15 23:20:16
 */
package todo_service

import (
	"context"
	"fmt"
	"nbox/src/model/dao/db"
	"time"
)

type TodoService interface {
}

type Todo struct {
	ID        uint64
	Content   string
	Status    uint8
	Creator   uint64
	Deleted   uint8
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type UpdateTodoParams struct {
	Content string
	Status  string
	Deleted string
}

func add(ctx context.Context, params *Todo) (*Todo, error) {
	todo := &db.Todo{
		ID:      params.ID,
		Content: params.Content,
		Status:  params.Status,
		Creator: params.Creator,
		Deleted: params.Deleted,
	}
	dbTodo, err := todo.AddTodo(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return &Todo{}, err
	}
	result := &Todo{
		ID:        dbTodo.ID,
		Content:   dbTodo.Content,
		Status:    dbTodo.Status,
		Creator:   dbTodo.Creator,
		Deleted:   dbTodo.Deleted,
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
		dbVals[db.TodoColumn.Deleted] = params.Deleted
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
