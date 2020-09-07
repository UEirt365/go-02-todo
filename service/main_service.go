package service

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ueirt/db"
	"ueirt/model"
)

func CreateTodo(context *gin.Context) {
	appGin := &model.Gin{C: context}

	var body model.TodoRequest
	if err := context.Bind(&body); err != nil {
		appGin.ToErrorResponse(http.StatusBadRequest, "Invalid request body!", err)
		return
	}

	val := validation.Validation{}
	ok, _ := val.Valid(&body)
	if !ok {
		appGin.ToErrorResponse(http.StatusBadRequest, "Invalid request body.", val.Errors)
		return
	}

	id, err := db.InsertNewTodo(body)
	if err != nil {
		appGin.ResponseFromError(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"id": id, "message": "Todo item created successfully!"})
}

func GetAllTodo(context *gin.Context) {
	todoList, err := db.SelectAllTodo()

	if err != nil {
		appGin := &model.Gin{C: context}
		appGin.ResponseFromError(err)
		return
	}

	context.JSON(http.StatusOK, todoList)
}

func GetTodoById(context *gin.Context) {
	todo, err := db.GetTodoById(context.Param("id"))

	if err != nil {
		appGin := &model.Gin{C: context}
		appGin.ResponseFromError(err)
		return
	}

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

	var todoRequest model.TodoRequest
	if err = context.Bind(&todoRequest); err != nil {
		appGin.ToErrorResponse(http.StatusBadRequest, "Invalid request body!", err)
		return
	}

	id, err := db.UpdateTodo(todoId, todoRequest)
	if err != nil {
		appGin.ResponseFromError(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": id, "message": "Todo item updated successfully!"})
}

func DeleteTodo(context *gin.Context) {
	appGin := &model.Gin{C: context}

	todoId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		appGin.ToErrorResponse(http.StatusBadRequest, "Invalid todo id", err)
		return
	}

	if err = db.DeleteTodo(todoId); err != nil {
		appGin.ResponseFromError(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo item was deleted!"})
}
