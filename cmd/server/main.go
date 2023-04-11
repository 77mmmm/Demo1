package main

import (
<<<<<<< HEAD
	"awesomeProject/handler"
	"awesomeProject/handler/middleware"
	"fmt"
	"net/http"

=======
	"fmt"
	"net/http"

	"awesomeProject/handler"
	"awesomeProject/handler/middleware"

>>>>>>> 6b8199d2ebeee8145588c5208812ddf694266766
	_ "github.com/go-sql-driver/mysql"
)

func main() {
<<<<<<< HEAD

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handler.Login)

	mux.HandleFunc("/commodity", handler.Product_Search)
	mux.HandleFunc("/purchase", handler.Purchase)

	mux.Handle("/hello", middleware.Auth(http.HandlerFunc(handler.Login)))
	fmt.Println("Server listening on port 8080...")
	_ = http.ListenAndServe("0.0.0.0:8080", middleware.Limit(mux))
=======
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/commodity", handler.Product_Search)
	http.HandleFunc("/purchase", handler.Purchase)

	http.Handle("/hello", middleware.Auth(http.HandlerFunc(handler.Login)))
	fmt.Println("Server listening on port 8080...")
	_ = http.ListenAndServe("0.0.0.0:8080", nil)
>>>>>>> 6b8199d2ebeee8145588c5208812ddf694266766
}
