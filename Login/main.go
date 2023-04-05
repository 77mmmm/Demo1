package main

import (
	"awesomeProject/Login/mysql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", mysql.Login)

	http.Handle("/hello", mysql.AuthMiddleware(http.HandlerFunc(mysql.Login)))
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
