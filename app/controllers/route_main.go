package controllers

import (
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
	_, err := session(w, r) //クッキーの取得
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
