package db

import (
	"github.com/go-gorp/gorp"
	"log"
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

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDb", "UTF8"}}
	return dbMap
}

func SelectAllTodo() []model.Todo {
	var todoList []model.Todo
	dbMap.Select(&todoList, "select * from todo")

	return todoList
}

func GetTodoById(id string) (model.Todo, bool) {
	var todo model.Todo
	err := dbMap.SelectOne(&todo, "select * from ueirt.todo where id = ?", id)
	if err != nil {
		return todo, false
	}

	return todo, true
}

func InsertNewTodo(todo model.Todo) int64 {
	//var todo dto.Todo
	insert, err := dbMap.Exec("INSERT INTO ueirt.todo (title, content) VALUES (?, ?)", todo.Title, todo.Content)
	if err != nil {
		log.Fatal("Failed to insert new todo to database", err)
	}
	id, _ := insert.LastInsertId()

	return id
}

func UpdateTodo(todo model.Todo) int64 {
	//var todo dto.Todo
	_, err := dbMap.Exec("UPDATE ueirt.todo set title = ?, content =? where id=?", todo.Title, todo.Content, todo.Id)
	if err != nil {
		log.Fatal("Failed to update new todo to database", err)
	}
	return int64(todo.Id)
}

func DeleteTodo(todoId int) bool {
	_, err := dbMap.Exec("DELETE FROM ueirt.todo where id=?", todoId)
	if err != nil {
		log.Fatal("Failed to delete todo", err)
		return false
	}

	return true
}
