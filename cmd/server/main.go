package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"time"

	"awesomeProject/handler"
	"awesomeProject/handler/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//组的概念
	//gin
	//url :   http://qq.com/user/login
	//        http://qq.com/user/register

	//分组
	//        http://qq.com/shoppingcart/list

	//restful Api    http  get  查   post  修改  put  创建  delete
	//  methaod: GET   http://qq.com/user/789

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handler.Login)

	mux.HandleFunc("/commodity", handler.Product_Search)
	mux.HandleFunc("/purchase", handler.Purchase)

	mux.Handle("/hello", middleware.Auth(http.HandlerFunc(handler.Login)))
	fmt.Println("Server listening on port 8080...")
	Limit := middleware.Limiter{
		Lm: rate.NewLimiter(rate.Every(time.Second), 10),
	}
	_ = http.ListenAndServe("0.0.0.0:8080", Limit.Limit(mux))

}
