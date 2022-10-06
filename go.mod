module todo_app

go 1.18

// go get -v ...アプリの中で使われている依存パッケージが表示される
require (
	github.com/google/uuid v1.3.0
	github.com/mattn/go-sqlite3 v1.14.15
	gopkg.in/go-ini/ini.v1 v1.67.0
)
