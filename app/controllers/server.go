package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top) //topというハンドラーの処理を実行(route_main.go)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}
