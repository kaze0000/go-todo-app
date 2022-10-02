package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...)) //Must...エラーだとpanicになる
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top) //topというハンドラーの処理を実行(route_main.go)
	http.HandleFunc("/signup", signup)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}
