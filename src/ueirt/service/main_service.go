package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ueirt/db"
	"ueirt/model"
)

func CreateTodo(context *gin.Context) {
	appGin := &model.Gin{C: context}

	var body model.Todo
	if err := context.Bind(&body); err != nil {
		appGin.ToErrorResponse(http.StatusBadRequest, "Invalid request body!", err)
		return
	}
	id := db.InsertNewTodo(body)
	context.JSON(http.StatusCreated, gin.H{"id": id, "message": "Todo item created successfully!"})
}

func GetAllTodo(context *gin.Context) {
	todoList := db.SelectAllTodo()
	context.JSON(http.StatusOK, todoList)
}

func GetTodoById(context *gin.Context) {
	todo := db.GetTodoById(context.Param("id"))
	context.JSON(http.StatusOK, todo)
}

func UpdateTodo(context *gin.Context) {
	appGin := &model.Gin{C: context}

	var todoId int
	var err error

	if todoId, err = strconv.Atoi(context.Param("id")); err != nil {
		appGin.ToErrorResponse(http.StatusBadRequest, "Invalid todo id", err)
		return
	}

	var body model.Todo
	if err = context.Bind(&body); err != nil {
		appGin.ToErrorResponse(http.StatusBadRequest, "Invalid request body!", err)
		return
	}

	body.Id = todoId
	id := db.UpdateTodo(body)

	context.JSON(http.StatusOK, gin.H{"id": id, "message": "Todo item updated successfully!"})
}

func DeleteTodo(context *gin.Context) {
	appGin := &model.Gin{C: context}

	todoId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		appGin.ToErrorResponse(http.StatusBadRequest, "Invalid todo id", err)
		return
	}

	db.DeleteTodo(todoId)

	context.JSON(http.StatusOK, gin.H{"message": "Todo item was deleted!"})
}
