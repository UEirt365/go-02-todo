package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"ueirt/db"
	"ueirt/dto"
)

func CreateTodo(context *gin.Context) {
	var body dto.Todo
	error := context.Bind(&body)
	if error != nil {
		log.Print("Fail to parse body", error)
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid request body!"})
		return
	}
	id := db.InsertNewTodo(body)
	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message": "Todo item created successfully!",
		"id":      id})
}

func GetAllTodo(context *gin.Context) {
	todoList := db.SelectAllTodo()

	context.JSON(http.StatusOK, todoList)
}

func GetTodoById(context *gin.Context) {
	todo, ok := db.GetTodoById(context.Param("id"))

	if !ok {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.JSON(http.StatusOK, todo)
}

func UpdateTodo(context *gin.Context) {
	todoId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid todo id"})
		return
	}

	var body dto.Todo
	error := context.Bind(&body)
	if error != nil {
		log.Print("Fail to parse body", error)
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid request body!"})
		return
	}
	body.Id = todoId
	id := db.UpdateTodo(body)

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
		"message": "Todo item updated successfully!",
		"id":      id})
}

func DeleteTodo(context *gin.Context) {
	todoId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid todo id"})
		return
	}

	ok := db.DeleteTodo(todoId)

	if ok {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo item was deleted!"})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Failed to delete todo!"})
	}
}
