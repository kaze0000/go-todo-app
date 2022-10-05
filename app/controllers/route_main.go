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

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r , "/todos", http.StatusFound)
	} else { //エラーがない。つまりログインしている場合
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err:= user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
