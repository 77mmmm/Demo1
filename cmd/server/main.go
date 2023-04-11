package main

import (
	"awesomeProject/handler"
	"awesomeProject/handler/middleware"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handler.Login)

	mux.HandleFunc("/commodity", handler.Product_Search)
	mux.HandleFunc("/purchase", handler.Purchase)

	mux.Handle("/hello", middleware.Auth(http.HandlerFunc(handler.Login)))
	fmt.Println("Server listening on port 8080...")
	_ = http.ListenAndServe("0.0.0.0:8080", middleware.Limit(mux))
}
