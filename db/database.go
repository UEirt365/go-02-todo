package db

import (
	"github.com/go-gorp/gorp"
	"log"
	"net/http"
	"ueirt/model"
)
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var dbMap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ueirt")

	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDb", Encoding: "UTF8"}}
	return dbMap
}

func SelectAllTodo() []model.Todo {
	var todoList []model.Todo

	if _, err := dbMap.Select(&todoList, "select * from todo"); err != nil {
		panic(model.ApiError{Status: http.StatusInternalServerError, Message: "Failed to select database", Err: err})
	}

	return todoList
}

func GetTodoById(id string) model.Todo {
	var todo model.Todo

	if err := dbMap.SelectOne(&todo, "select * from ueirt.todo where id = ?", id); err != nil {
		panic(model.ApiError{Status: http.StatusNotFound, Message: "Todo not found", Err: err})
	}

	return todo
}

func InsertNewTodo(todo model.Todo) int64 {
	insert, err := dbMap.Exec("INSERT INTO ueirt.todo (title, content) VALUES (?, ?)", todo.Title, todo.Content)

	if err != nil {
		panic(model.ApiError{Status: http.StatusInternalServerError, Message: "Failed to insert database", Err: err})
	}

	id, err := insert.LastInsertId()
	if err != nil {
		panic(model.ApiError{Status: http.StatusInternalServerError, Message: "Failed to get LastInsertId", Err: err})
	}

	return id
}

func UpdateTodo(todo model.Todo) int64 {
	result, err := dbMap.Exec("UPDATE ueirt.todo set title = ?, content =? where id=?", todo.Title, todo.Content, todo.Id)
	if err != nil {
		panic(model.ApiError{Status: http.StatusInternalServerError, Message: "Failed to update todo to database", Err: err})
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected <= 0 {
		panic(model.ApiError{Status: http.StatusNotFound, Message: "Todo not found.", Err: err})
	}
	if err != nil {
		panic(model.ApiError{Status: http.StatusInternalServerError, Message: "Failed to get updated todo", Err: err})
	}

	return int64(todo.Id)
}

func DeleteTodo(todoId int) {
	result, err := dbMap.Exec("DELETE FROM ueirt.todo where id=?", todoId)
	if err != nil {
		panic(model.ApiError{Status: http.StatusInternalServerError, Message: "Failed to delete todo", Err: err})
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected <= 0 {
		panic(model.ApiError{Status: http.StatusNotFound, Message: "Todo not found.", Err: err})
	}
	if err != nil {
		panic(model.ApiError{Status: http.StatusInternalServerError, Message: "Failed to get updated todo", Err: err})
	}
}
