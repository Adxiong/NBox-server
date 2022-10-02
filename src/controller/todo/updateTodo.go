/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-09-25 22:49:26
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-02 23:27:25
 */
package todo

import (
	"log"
	todo_service "nbox/src/model/service/todo"

	"github.com/gin-gonic/gin"
)

type TodoParams struct {
	TID     uint64 `json:"tid" binding:"required"`
	Content string `json:"content"`
	Status  uint8  `json:"status"`
}

func UpdateTodo(ctx *gin.Context) {
	params, errParams := checkUpdateTodoParams(ctx)

	if errParams != nil {
		log.Println("error", errParams)
		ctx.JSON(200, "msg:params is invalid")
		return
	}

	todoService := todo_service.NewTodoService()
	todoParams := &todo_service.UpdateTodoParams{
		Content: params.Content,
		Status:  params.Status,
	}
	rowsAffected, errUpdate := todoService.UpdateByTID(ctx, params.TID, todoParams)

	if errUpdate != nil {
		log.Println("error", errUpdate)
		ctx.JSON(200, "msg: update failed")
	}

	ctx.JSON(200, "ok")
}

func checkUpdateTodoParams(ctx *gin.Context) (*TodoParams, error) {
	todoParams := &TodoParams{}
	err := ctx.BindJSON(todoParams)

	if err != nil {
		log.Println("error", err)
		return todoParams, err
	}

	return todoParams, nil
}
