package handler

import (
	"awesomeProject/logic/user"
	"fmt"
	"net/http"

	"awesomeProject/lib/ctx"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	user := user.RegisterUser{Ctx: ctx.GetGlobalCtx()}
	Bool, err := user.Register(username, password, email)

	if err != nil {
		http.Error(w, fmt.Sprintf("error checking user: %s", err), http.StatusBadRequest)
	}
	if Bool == false {
		http.Error(w, "注册失败", http.StatusBadRequest)
<<<<<<< HEAD
		fmt.Println("注册失败")

	} else {
		fmt.Println("注册成功")
		fmt.Println("name:", username, "pass:", password, "email", email)
=======

	} else {
		fmt.Println("注册成功")
>>>>>>> 6b8199d2ebeee8145588c5208812ddf694266766
	}

}
