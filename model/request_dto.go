package model

type Todo struct {
	Id      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}

type TodoRequest struct {
	Title   string `valid:"Required"`
	Content string `valid:"Required"`
}
