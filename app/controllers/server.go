package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
	"todo_app/app/models"
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

func session(_ http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc { //handlerを関数を返す関数
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1のようなURLからidを受けとって処理する
		q := validPath.FindStringSubmatch(r.URL.Path) // マッチした部分をスライスで取得する
		if q == nil { //何もマッチしない場合
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2]) //最後のパスをint型として受け取る
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi) //引数で渡したtodoEdit()を実行
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top) //topというハンドラーの処理を実行(route_main.go)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	return http.ListenAndServe(":" + config.Config.Port, nil)
}
