package controllers

import (
	"context"
	"log"

	"github.com/chirani/templ_go_project/db"
)

type Todo struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	IsDone  uint8  `json:"is_done"`
}

func CreateTodo(title string, details string) error {

	statement := `insert into todos(title, details, is_done) values($1, $2, $3);`

	_, err := db.DB.Exec(statement, title, details, 0)

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func GetAllTodos() ([]Todo, error) {
	todos := []Todo{}
	statement := `SELECT id, title, details, is_done FROM todos;`
	rows, err := db.DB.Query(statement)

	if err != nil {

		return todos, err
	}
	var (
		title   string
		text    string
		is_done uint8
		id      uint64
	)

	for rows.Next() {

		err := rows.Scan(&id, &title, &text, &is_done)
		if err != nil {
			return todos, err
		}

		todo := Todo{
			Id:      id,
			Title:   title,
			IsDone:  is_done,
			Details: text,
		}

		todos = append(todos, todo)
	}

	return todos, err
}

func GetTodo(ctx context.Context, id uint64) (Todo, error) {
	statement := "SELECT * from todos WHERE id=$1"
	todo := Todo{}

	row := db.DB.QueryRowContext(ctx, statement, id)

	var (
		title   string
		text    string
		is_done uint8
	)

	err := row.Scan(title, text, is_done)
	if err != nil {
		return todo, err
	}

	todo.Id = id
	todo.IsDone = is_done
	todo.Details = text
	todo.Title = title

	return todo, err
}

func DeleteTodos(id uint64) error {
	statement := `DELETE FROM todos WHERE id=$1`
	_, err := db.DB.Exec(statement, id)

	if err != nil {
		return err
	}
	return nil
}

func MarkDone(ctx context.Context, id uint64) error {
	todo, err := GetTodo(ctx, id)
	if err != nil {
		return err
	}
	statement := `UPDATE todos SET is_done=$1 WHERE id=$2`

	if todo.IsDone == 1 {
		todo.IsDone = 0
	} else {
		todo.IsDone = 1
	}

	_, err = db.DB.QueryContext(ctx, statement, todo.IsDone, id)
	if err != nil {
		return err
	}
	return nil
}
