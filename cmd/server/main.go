package main

import (
	"fmt"
	"net/http"

	"awesomeProject/handler"
	"awesomeProject/handler/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/commodity", handler.Product_Search)
	http.HandleFunc("/purchase", handler.Purchase)

	http.Handle("/hello", middleware.Auth(http.HandlerFunc(handler.Login)))
	fmt.Println("Server listening on port 8080...")
	_ = http.ListenAndServe("0.0.0.0:8080", nil)
}
