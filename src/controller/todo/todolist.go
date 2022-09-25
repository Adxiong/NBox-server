/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-09-25 22:33:01
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-09-25 22:45:47
 */
package todo

import (
	"log"
	todo_service "nbox/src/model/service/todo"

	"github.com/gin-gonic/gin"
)

type ReqTodoListParams struct {
	Page uint64 `form:"page"`
	Size uint64 `form:"size"`
}

func TodoList(ctx *gin.Context) {
	params, errParams := checkReqTodoListParams(ctx)
	if errParams != nil {
		log.Println("error", errParams)
		ctx.JSON(200, gin.H{"msg": "params is invalid"})
		return
	}

	var UID uint64 = 11111

	todoService := todo_service.NewTodoService()

	list, errList := todoService.FindListByUID(ctx, UID, params.Page, params.Size)
	if errList != nil {
		log.Println("error", errList)
		ctx.JSON(200, gin.H{"msg": "get todoList failed"})
		return
	}

	ctx.JSON(200, gin.H{"msg": "success", "data": *list})

}

func checkReqTodoListParams(ctx *gin.Context) (*ReqTodoListParams, error) {
	params := &ReqTodoListParams{}
	err := ctx.ShouldBind(params)

	if err != nil {
		log.Println("error", err)
		return params, err
	}
	return params, nil
}
