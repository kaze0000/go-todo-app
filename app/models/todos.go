package models

import (
	"fmt"
	"log"
	"time"
)

type Todo struct {
	ID int
	Content string
	UserID int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `select id, content, user_id, created_at from todos
	where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select * from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
										&todo.Content,
										&todo.UserID,
										&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (u *User) GetoTodoByUser() (todos []Todo, err error) {
	cmd := `select * from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
										&todo.Content,
										&todo.UserID,
										&todo.CreatedAt,)
		if err != nil {
			fmt.Println(err) //log.Fatallnはos.Exitしていて、アプリケーションが終わってしまうので、あまり使わない
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}
