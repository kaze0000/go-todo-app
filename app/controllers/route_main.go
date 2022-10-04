package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) { //ハンドラー
	/*
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, "hello")
	*/
	_, err := session(w, r)
	if err != nil { //ログインしていない
		generateHTML(w, "hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r) //クッキーと一致するsessionの取得
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetoTodoByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index") //第２引数userはtemplateにわたすやつ
	}
}
