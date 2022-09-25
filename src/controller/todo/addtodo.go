/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-09-25 22:16:43
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-25 22:31:47
 */
package todo

import (
	"log"
	todo_service "nbox/src/model/service/todo"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Content string `form:"content" binding:"required"`
	Status  uint8  `form:"status" binding:"status"`
}

func AddTodo(ctx *gin.Context) {
	params, errParams := checkAddTodoParams(ctx)
	if errParams != nil {
		log.Println("errror", errParams)
		ctx.JSON(200, "msg: params is invalid")
		return
	}

	todoService := todo_service.NewTodoService()
	todoParams := &todo_service.Todo{
		Content: params.Content,
		Status:  params.Status,
	}
	todo, errTodo := todoService.Add(ctx, todoParams)

	if errTodo != nil {
		log.Println("error", errTodo)
		ctx.JSON(200, gin.H{"msg": "AddTodo Failed"})
		return
	}

	ctx.JSON(200, gin.H{"msg": "success", "data": *todo})
}

func checkAddTodoParams(ctx *gin.Context) (*Todo, error) {
	params := &Todo{}
	err := ctx.ShouldBind(params)
	if err != nil {
		log.Println("error", err)
		return params, err
	}
	return params, nil
}
