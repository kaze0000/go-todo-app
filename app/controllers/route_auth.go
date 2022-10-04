package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil { //ログインしていない
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm() //フォームの値を取得
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name: r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", http.StatusFound) //302リダイレクト
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil { //ログインしていない
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
 		log.Println(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie {
			Name: "_cookie",
			Value: session.UUID,
			HttpOnly: true, //Cookie属性としてこれを付与するとJavaScriptからアクセスできなくなる。
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", http.StatusFound)
	} else { //パスワードが一致しない場合
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value} //構造体を作って、UUIDに取得したクッキーのvalueを入れる
		session.DeleteSessionByUUID()
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
