/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-15 00:37:26
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-03 00:04:41
 */
package todo

import (
	"log"
	todo_service "nbox/src/model/service/todo"

	"github.com/gin-gonic/gin"
)

type DelTodoParams struct {
	TID uint64 `json:"tid" binding:"required"`
}

func DeleteTodo(ctx *gin.Context) {
	params, errParams := checkDelTodoParams(ctx)

	if errParams != nil {
		log.Println("error", errParams)
		ctx.JSON(200, "msg:params is invalid")
		return
	}

	todoService := todo_service.NewTodoService()
	rowsAffected, errUpdate := todoService.DeleteByTID(ctx, params.TID)

	if errUpdate != nil {
		log.Println("error", errUpdate)
		ctx.JSON(200, "msg: delete failed")
	}

	if rowsAffected == 0 {
		ctx.JSON(200, "msg: delete failed")
	}

	ctx.JSON(200, "ok")
}

func checkDelTodoParams(ctx *gin.Context) (*DelTodoParams, error) {
	todoParams := &DelTodoParams{}
	err := ctx.BindJSON(todoParams)

	if err != nil {
		log.Println("error", err)
		return todoParams, err
	}

	return todoParams, nil
}
