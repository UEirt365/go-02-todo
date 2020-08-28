package main

import (
	"github.com/gin-gonic/gin"
	"ueirt/service"
)

func main() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.POST("", service.CreateTodo)
		v1.GET("", service.GetAllTodo)
		v1.GET(":id", service.GetTodoById)
		v1.PUT(":id", service.UpdateTodo)
		v1.DELETE(":id", service.DeleteTodo)
	}
	router.Run(":8080")
}
