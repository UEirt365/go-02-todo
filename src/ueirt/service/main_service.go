package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ueirt/db"
	"ueirt/model"
)

func CreateTodo(context *gin.Context) {
	defer HandleError(context)
	var body model.Todo
	if err := context.Bind(&body); err != nil {
		panic(model.ApiError{Status: http.StatusBadRequest, Message: "Invalid request body!", Err: err})
	}
	id := db.InsertNewTodo(body)
	context.JSON(http.StatusCreated, gin.H{"id": id, "message": "Todo item created successfully!"})
}

func GetAllTodo(context *gin.Context) {
	defer HandleError(context)
	todoList := db.SelectAllTodo()
	context.JSON(http.StatusOK, todoList)
}

func GetTodoById(context *gin.Context) {
	defer HandleError(context)
	todo := db.GetTodoById(context.Param("id"))
	context.JSON(http.StatusOK, todo)
}

func UpdateTodo(context *gin.Context) {
	defer HandleError(context)

	var todoId int
	var err error

	if todoId, err = strconv.Atoi(context.Param("id")); err != nil {
		panic(model.ApiError{Status: http.StatusBadRequest, Message: "Invalid todo id", Err: err})
	}

	var body model.Todo
	if err = context.Bind(&body); err != nil {
		panic(model.ApiError{Status: http.StatusBadRequest, Message: "Invalid request body!", Err: err})
	}

	body.Id = todoId
	id := db.UpdateTodo(body)

	context.JSON(http.StatusOK, gin.H{"id": id, "message": "Todo item updated successfully!"})
}

func DeleteTodo(context *gin.Context) {
	defer HandleError(context)
	todoId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic(model.ApiError{Status: http.StatusBadRequest, Message: "Invalid todo id", Err: err})
	}

	db.DeleteTodo(todoId)

	context.JSON(http.StatusOK, gin.H{"message": "Todo item was deleted!"})
}

func HandleError(context *gin.Context) {
	if lastError := recover(); lastError != nil {
		switch lastError.(type) {
		case model.ApiError:
			apiError := lastError.(model.ApiError)
			context.JSON(apiError.ToErrorResponse())
		default:
			context.JSON(http.StatusInternalServerError, lastError)
		}
	}
}
