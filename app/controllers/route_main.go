package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, _  *http.Request) { //ハンドラー
	/*
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, "hello")
	*/
	generateHTML(w, "hello", "layout", "public_navbar","top")
}
