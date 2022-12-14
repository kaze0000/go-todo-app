package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main(){
	/*
	fmt.Println(config.Config.Port) //configパッケージのConfigという変数のPortフィールド
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
	*/

	fmt.Println(models.Db) //テーブルの作成

	/*ユーザーの作成
	u := &models.User{}
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.Password = "testtest"
	log.Println(u)

	u.CreateUser()
	*/

	/*
	u, _ := models.GetUser(2)
	log.Println(u)
	*/

	/*
	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	log.Println(u)
	*/

	/*
	u.DeleteUser()
	u, _ = models.GetUser(2)
	log.Println(u)
	*/

	/*
	user, _ := models.GetUser(2)
	user.CreateTodo("First Todo")
	*/

	/*
	t, _ := models.GetTodo(1)
	log.Println(t)
	*/

	/*
	user, _ := models.GetUser(3)
	user.CreateTodo("Third Todo")
	*/

	/*
	todos, _ := models.GetTodos()
	for _, v := range todos {
		log.Println(v)
	}
	*/

	/*
	user2, _ := models.GetUser(2)
	todos, _ := user2.GetoTodoByUser()

	for _, v := range todos {
		log.Println(v)
	}

	t, _ := models.GetTodo(1)
	t.Content = "Update Todo"
	t.UpdateTodo()
	*/

	controllers.StartMainServer()

	/*
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

	valid, _ := session.CheckSession()
	fmt.Println(valid)
	*/
}
